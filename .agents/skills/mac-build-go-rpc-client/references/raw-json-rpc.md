# Raw JSON-RPC over HTTP

Use this variant only when the user chooses direct protocol implementation or a dependency review justifies avoiding an SDK.

## Transport

- Inject `*http.Client` or a narrow `Do(*http.Request)` interface.
- Require caller context on every request and define transport-level timeout policy deliberately.
- Encode typed JSON-RPC request and response envelopes with `encoding/json`.
- Generate concurrency-safe request IDs and validate response IDs when protocol correctness depends on correlation.
- Set expected content headers and bound response body reads.
- Distinguish a missing result, JSON `null`, and a zero-valued decoded object when the API gives them different meanings.
- Preserve typed JSON-RPC error code, message, and safe structured data.

## Failure handling

- Classify network errors, HTTP statuses, malformed envelopes, and JSON-RPC errors separately.
- A `2xx` response can still carry a JSON-RPC error. A non-`2xx` response may carry useful structured error data.
- Treat encoding, invalid endpoint construction, schema mismatch, and unsupported protocol versions as permanent.
- Apply provider-specific handling for rate limiting and server-busy codes.
- Do not retry mutations without an idempotency guarantee.

## Scope control

Implement only the required protocol surface. Do not add batching, notifications, WebSocket support, middleware stacks, or generic reflection-based method generation until a current consumer requires them.

## Tests

Use `httptest.Server` to assert request envelopes, parameters, headers, response ID handling, JSON-RPC error decoding, HTTP classification, retry behavior, body limits, malformed responses, and cancellation. Keep fixtures minimal but valid for the actual decoder.
