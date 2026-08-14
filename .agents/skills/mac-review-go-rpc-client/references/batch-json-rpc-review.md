# JSON-RPC Batch Review

Apply this checklist only when the implementation sends or consumes JSON-RPC batches.

## Contract and mapping

- Confirm batching addresses an actual requirement and the public contract documents empty input, duplicate inputs, ordering, limits/chunking, and partial success.
- Trace each input index to its `rpc.BatchElem`, typed result destination, per-item `BatchElem.Error`, decoded value, and final result. Flag any path that can misalign values, drop duplicates, or report false success.
- Confirm a nil `BatchCallContext` error is not treated as success without inspecting every item error; confirm a non-nil transport error is applied consistently to affected pending items.
- Check null, missing, malformed, and decode-failed item results cannot panic or be mistaken for valid zero values.

## Retry, context, and concurrency

- Verify only transient, idempotent failed items re-enter a later sub-batch after a successful response. Successful and permanent-failure items must not be repeated.
- For a transport-level failure, verify retrying the entire pending subset is safe; mutation batches require explicit idempotency/deduplication evidence.
- Verify each retry copies results/errors correctly back to original positions and preserves final error identity after exhaustion.
- Check cancellation/deadline stops sends and backoff waits immediately. Confirm `MaxAttempts` means total batch sends and no retry occurs after the final attempt.
- Inspect shared slices, result pointers, counters, random sources, and `Close` behavior for races when concurrent use is promised.

## Tests and logs

- Look for deterministic fake-batcher tests covering partial retry, permanent no-retry, transport retry, exhaustion, cancellation, empty input, order/duplicates, and aggregate result semantics.
- Look for a protocol-level test that asserts a JSON array request, correct method/params, item error envelopes, and out-of-order response correlation.
- Ensure logs carry safe operation/attempt/pending/final-count fields but never full batch payloads, secrets, or raw transaction data.
