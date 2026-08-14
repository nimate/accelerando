# Quality and Testing

## Boundary checklist

- Start from required operations and outputs, not a preferred directory tree.
- Keep interfaces narrow and owned by the consumer or wrapper that needs substitution.
- Keep domain models free of third-party types unless SDK interoperability is the explicit API goal.
- Keep the production constructor separate from the dependency-injected constructor used by unit tests.
- Validate nil loggers, transports, clocks, backoff factories, endpoints, and invalid retry values.
- Provide lifecycle cleanup when dialing creates connections, goroutines, or idle transports.
- Document whether the client and `Close` are safe for concurrent use.
- Keep credentials out of errors, logs, metrics labels, and endpoint fields.

## Unit tests

Cover behavior rather than implementation structure:

- successful call and exact domain mapping,
- exact method name and arguments passed to the dependency,
- immediate success uses one attempt,
- transient failures eventually succeed,
- retry exhaustion returns the final meaningful error,
- permanent errors use one attempt,
- canceled and expired contexts stop promptly,
- nil or malformed successful responses do not panic,
- injected clock controls fetched timestamps,
- `Close` delegates exactly as promised.

Use a small fake implementing the narrow interface. Inject zero-delay backoff or a fake sleeper instead of sleeping in tests.

## Transport tests

Use `httptest.Server` or an in-process server to verify behavior owned by the adapter:

- request method, envelope, headers, and parameter encoding,
- response and error-envelope decoding,
- HTTP status classification,
- response size limits and malformed bodies,
- cancellation while a request is in flight,
- retry hints when supported.

Do not use a public endpoint in the default test suite. Put live compatibility checks behind an explicit environment variable or build tag.

## Advanced batch tests

When batching is required, additionally cover:

- empty input makes no request and duplicate inputs retain their positions,
- all successes and mixed per-item successes/errors are mapped to the matching input,
- permanent item errors are not retried while only transient items are retried,
- a transient whole-batch transport failure retries only safe pending work,
- retry exhaustion preserves each item’s final error and the aggregate outcome is documented,
- out-of-order responses, null results, malformed item data, and configured item/payload limits,
- cancellation during a batch call and during the retry wait.

## Completion gate

- No blanket retry of every error.
- No production delays in unit tests.
- No SDK type leakage by accident.
- No missing cleanup for owned resources.
- No claims of success without formatter, build, and tests.
- Run the race detector when the client contains shared mutable state or is advertised as concurrent-safe.
