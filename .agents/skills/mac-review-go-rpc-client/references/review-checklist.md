# RPC Client Review Checklist

## Contract and boundary

- Confirm each public method has a real consumer and stable semantics.
- Check whether third-party SDK types escape unintentionally into domain APIs.
- Prefer narrow consumer-owned interfaces; do not demand an interface where no substitution boundary exists.
- Check constructors for invalid endpoints, nil dependencies, unusable retry values, and partial initialization.
- Verify domain mapping preserves units, timestamps, identifiers, optional values, and protocol-specific null semantics.

## Context and lifecycle

- Propagate caller context through dialing, requests, retries, and response processing.
- Stop backoff waits promptly on cancellation or deadline expiry.
- Check whether internal timeouts conflict with or silently override caller deadlines.
- Close owned connections, response bodies, subscriptions, timers, and goroutines on every exit path.
- Verify `Close` semantics, repeated calls, and behavior during concurrent operations when relevant.

## Errors and retries

- Separate local validation, encoding, network, HTTP, protocol, decoding, and domain errors.
- Preserve error identity for `errors.Is` and `errors.As`.
- Retry only classified transient failures for operations safe to repeat.
- Treat mutation retries as unsafe without idempotency or deduplication guarantees.
- Confirm whether configuration counts total attempts or retries after the initial call.
- Check exponential growth, jitter, maximum delay, elapsed budget, and fresh per-call backoff state.
- Honor documented retry hints with a defensive cap.
- Avoid logging both every attempt and the same final error at multiple layers without purpose.

## Transport correctness

- Validate method names, parameters, headers, IDs, and response correlation.
- Bound response bodies and handle malformed or unexpectedly successful responses.
- Do not assume HTTP success means RPC success or HTTP failure lacks structured error data.
- Distinguish a missing value, JSON `null`, and a valid zero value when required by the protocol.
- Check request ID generation and all shared mutable fields for races.

## Observability and security

- Log operation and safe endpoint identity, attempts, delays, and final outcome using structured fields.
- Redact credentials in URLs, headers, request bodies, error strings, and metrics labels.
- Avoid high-cardinality labels and unbounded payload logging.
- Confirm logs distinguish retry exhaustion, permanent failure, and caller cancellation.
- Check whether metrics or traces are expected by existing repository conventions before requiring them.

## Scope discipline

- Flag missing batching, caching, fallback endpoints, subscriptions, or circuit breakers only when a current requirement needs them.
- Prefer a focused fix over a generic client framework.
- Treat style-only differences as non-findings unless they obscure a concrete defect.
