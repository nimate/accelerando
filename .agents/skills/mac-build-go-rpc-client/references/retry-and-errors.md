# Retry and Error Classification

## Define attempts precisely

Prefer `MaxAttempts` when configuration is user-facing. `MaxAttempts: 4` means one initial call plus at most three retries. If a library accepts `MaxRetries`, translate and document the difference at the boundary.

Use bounded exponential backoff with jitter:

- explicit initial delay,
- multiplier greater than one,
- maximum delay,
- maximum attempts or elapsed-time budget,
- cancellation-aware waiting.

Avoid global randomness when deterministic tests or high concurrency make an injected jitter source more appropriate.

## Classify before retrying

Typical retry candidates for an idempotent operation include temporary network failures, HTTP `408`, `425`, `429`, selected `5xx` responses, and documented server-busy RPC errors. Classification remains protocol- and provider-specific.

Typical permanent failures include invalid input, encoding failures, authentication and authorization failures, unsupported methods, invalid parameters, schema mismatch, and malformed successful responses. Treat `context.Canceled` and `context.DeadlineExceeded` as terminal for the current call.

Do not assume every JSON-RPC error is permanent or every HTTP `5xx` response is safe to retry. Capture the provider contract in a classifier function and test its table of cases.

## Preserve useful errors

- Wrap with operation context using `%w`.
- Preserve typed transport or RPC errors for `errors.As`.
- Return the caller context error when cancellation ends a backoff wait.
- Avoid replacing the root error with a generic "retry failed" message.
- Include attempt count in logs or metrics, not by destroying error identity.

## Respect server guidance

Parse `Retry-After` for HTTP APIs that define it. Apply configured caps so a malicious or broken server cannot force an unbounded wait. Record the selected delay without exposing request secrets.

## Mutating calls

Do not retry a mutating call merely because the transport failed before a response arrived; the server may already have committed it. Retry only when the protocol supplies a stable idempotency key, a deduplication guarantee, or operation-specific semantics that make repetition safe. Test ambiguous outcomes explicitly.
