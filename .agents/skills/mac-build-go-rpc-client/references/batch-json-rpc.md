# Advanced JSON-RPC Batch Client

Use this variant only for a current need to execute multiple independent, retry-safe requests. A batch is an optimization of concurrent reads, not a replacement for a clear single-call API.

## Contract

- Keep input/output alignment: `results[i]` always describes `inputs[i]`, including duplicates and failed items.
- Define partial success explicitly. Prefer a result type containing the requested identity, decoded value, and per-item error; return an aggregate error only if callers also need a whole-operation signal.
- Treat an empty input as a no-op without a network call.
- State and test maximum items and payload-size behavior. Chunk only when provider/node limits or a consumer requirement make it necessary; preserve original ordering across chunks.

## go-ethereum implementation

- Depend on a narrow `BatchCaller` containing `BatchCallContext(context.Context, []rpc.BatchElem) error`; production wiring may own `*rpc.Client` and expose `Close`.
- Build `rpc.BatchElem` values with a typed `Result` destination for every item. `BatchCallContext` correlates responses; map each `BatchElem.Error` and decoded result back to its original input index.
- A non-nil return from `BatchCallContext` is a batch/transport-level failure. A nil return does not mean every item succeeded: inspect every `BatchElem.Error`.
- Do not assume response order, manufacture request IDs, or use a shared mutable result buffer. Preserve result pointers correctly when retrying sub-batches.

## Retry and observability

- Retry only failed, transient, idempotent items after a successful batch response; leave successful and permanent-error items out of later attempts.
- For a transient transport-level failure, retry the affected sub-batch only when repeating all of its operations is safe. Preserve the final error on every affected item after exhaustion.
- Abort immediately for caller cancellation/deadline; make backoff waits cancellation-aware. Define `MaxAttempts` as total sends, including the first one.
- Log operation, attempt, pending count, final success/failure counts, and a redacted endpoint. Do not log request arguments, full responses, credentials, or raw transaction data by default.

## Tests

- Use a fake `BatchCaller` to assert first-batch composition, retry sub-batch composition, per-item transient recovery, permanent no-retry, transport retry, exhaustion, order/duplicate preservation, and cancellation during wait.
- Add an `httptest` or in-process JSON-RPC test for the actual batch array, method/parameter encoding, out-of-order responses, item error envelopes, null/malformed item results, and provider batch limits if supported.
