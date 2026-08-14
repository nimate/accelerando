---
name: mac-build-go-rpc-client
description: Design, implement, or extend production-quality RPC clients in Go, including Ethereum JSON-RPC with go-ethereum. Use when a task needs a basic single-request client or an explicit advanced JSON-RPC batch client, with narrow interfaces, domain mapping, context cancellation, bounded retry/backoff and jitter, structured logging, lifecycle cleanup, safe Ethereum hex handling, and deterministic tests. Do not use for a review-only request that does not ask for implementation.
---

# Build Go RPC Client

Build the smallest client that satisfies the consumer while making network behavior explicit, observable, and testable. Adapt the workflow to the repository instead of imposing a fixed package layout.

## Choose the capability level

Implement the **basic** variant by default: one logical request maps to one RPC request. Implement the **advanced batch** variant only when the consumer needs multiple independent read-only calls, lower round-trip overhead, or partial per-item outcomes. Do not add batching speculatively.

- For the basic variant, use the normal context-aware call exposed by the chosen adapter (`ethclient` method or `rpc.Client.CallContext`). Keep retries and one final result at the operation boundary.
- For the batch variant, read [batch-json-rpc.md](references/batch-json-rpc.md) before designing or coding it. With go-ethereum use `rpc.Client.BatchCallContext` and `rpc.BatchElem`; do not hand-roll JSON-RPC IDs, envelopes, response matching, or batch decoding.

## Establish the contract

Before editing code:

1. Inspect repository instructions, `go.mod`, existing transports, logging, configuration, domain types, and test conventions.
2. Identify the RPC protocol, required methods, inputs, domain outputs, endpoint/authentication requirements, and expected call volume.
3. Establish whether each operation is safe to retry. Treat mutating calls as non-idempotent until the API contract proves otherwise.
4. Determine the transport strategy:
   - named library or SDK,
   - raw protocol using the Go standard library,
   - comparison followed by a recommendation.

If the transport strategy is neither stated nor confidently established by the repository, stop and ask one concise question:

> Should the client use an existing library or SDK? If so, which one? Or should it implement the protocol directly with the Go standard library? I can compare both approaches first.

Do not ask this when the request already names a library such as `go-ethereum`, or existing code clearly establishes the choice. Ask about other missing requirements only when implementing a reasonable default would create material behavioral or compatibility risk.

## Plan the boundary

1. Define behavior before directories and abstractions.
2. Keep consumer-facing domain types independent of transport and SDK types unless the user explicitly wants an SDK-shaped API.
3. Define narrow interfaces at the consuming boundary. Include only methods used now.
4. Return concrete structs from constructors. Accept interfaces where substituting I/O behavior improves tests.
5. Keep dialing and concrete SDK/HTTP wiring in an adapter or constructor. Provide `Close` when the underlying transport owns resources.
6. Inject nondeterministic dependencies used by core logic, such as time, sleep/backoff, request IDs, or transport calls.

Read [quality-and-testing.md](references/quality-and-testing.md) before implementing the boundary and tests.

## Design failure behavior

Define an error matrix before adding retry:

- Retry only transient failures for operations known to be safe to repeat.
- Return validation, authentication, authorization, malformed response, unsupported method, and caller cancellation errors without retry.
- Preserve error identity with wrapping so callers can use `errors.Is` and `errors.As`.
- Honor `context.Canceled` and `context.DeadlineExceeded` immediately.
- Use bounded exponential backoff with jitter, an explicit attempt definition, and a cap.
- Honor server retry hints such as `Retry-After` when applicable.
- Create fresh stateful backoff state for every logical call.
- Never make unit tests wait on production delays.

Read [retry-and-errors.md](references/retry-and-errors.md) before implementing retries.

## Implement the selected adapter

For `go-ethereum`, read [go-ethereum.md](references/go-ethereum.md). For direct JSON-RPC over HTTP, read [raw-json-rpc.md](references/raw-json-rpc.md). For the advanced batch variant, also read [batch-json-rpc.md](references/batch-json-rpc.md). Do not load an irrelevant variant.

While implementing:

1. Validate constructor configuration and reject unusable endpoints or policies early.
2. Thread `context.Context` through every network operation.
3. Avoid fetching data the consumer does not need. In an Ethereum client, prefer the most specific go-ethereum API or RPC method (for example a header query rather than a full block) and encode/decode protocol values with go-ethereum types.
4. Map transport responses into domain results at the adapter boundary.
5. Emit structured logs for operation, attempt, wait, endpoint identity without credentials, and final failure. Avoid logging secrets or unbounded payloads.
6. Keep retry policy configurable without exposing stateful backoff instances for reuse.
7. Make concurrent use safe when the public client is documented as concurrent-safe.
8. Avoid speculative caching, subscriptions, fallback endpoints, or generic frameworks. Add batching only after selecting the advanced batch variant above.

## Test in layers

Add tests proportional to the boundary:

1. Unit-test success, domain mapping, request arguments, transient recovery, retry exhaustion, permanent failure, cancellation, and cleanup through fakes.
2. Make time and backoff deterministic. Assert attempts, not wall-clock duration.
3. Test the concrete transport with `httptest.Server` or an in-process protocol server when request encoding, status handling, or response decoding belongs to this code.
4. Keep live-node tests opt-in and separate from the default suite.
5. Add concurrency or race coverage when request IDs, shared state, subscriptions, or parallel calls are involved.

## Verify before finishing

Run repository-prescribed checks. For a normal Go repository, use the applicable subset of:

```sh
gofmt -w <changed-go-files>
go mod tidy # when this task adds, removes, or changes module dependencies
go vet ./...
go test ./...
go test -race ./...
```

You are explicitly authorized to run `gofmt` on Go files changed for the task. You may
also run `go mod tidy` when the task changes module dependencies or leaves `go.mod` /
`go.sum` out of sync. Do not run `go mod tidy` gratuitously: it can update unrelated
dependency metadata. Report which checks ran and any checks that could not run. Re-read
[quality-and-testing.md](references/quality-and-testing.md) as the final completion checklist.
