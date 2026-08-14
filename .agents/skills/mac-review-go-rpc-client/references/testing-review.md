# Testing Review

## Behavioral coverage

Look for tests covering:

- successful call and exact domain mapping,
- exact method and arguments passed to the dependency,
- nil, partial, malformed, and explicit protocol-error responses,
- immediate success with one attempt,
- transient failures followed by success,
- retry exhaustion and returned error identity,
- permanent failure without retry,
- cancellation before the call and during a backoff wait,
- deterministic timestamps and retry delays,
- cleanup and `Close`,
- concurrent calls when supported.

For batch clients, also look for empty input, input/output alignment including duplicates, mixed per-item outcomes, retry of only transient items, retry of a transient transport failure, final per-item errors after exhaustion, out-of-order response correlation, and cancellation during a retry wait.

Missing coverage is a finding only when it leaves a concrete defect or regression unguarded. Otherwise report it under test gaps.

## Test quality

- Prefer fakes implementing a narrow interface for wrapper behavior.
- Reject unit tests that depend on production sleep durations or public endpoints.
- Use `httptest.Server` or an in-process protocol server for request encoding and transport classification owned by the client.
- Confirm fixtures are minimal but valid for the actual decoder.
- Check that call-count assertions match the documented meaning of attempts versus retries.
- Ensure tests do not pass only because shared global state or ordering leaks between cases.

## Verification strategy

Run the narrow package test first when investigating a finding, then the repository suite. Use `go test -race` for shared request IDs, mutable retry state, subscriptions, caches, or advertised concurrent safety. Keep live compatibility tests opt-in through a build tag or explicit environment variable.

When a suspected issue cannot be reproduced, explain the remaining evidence and reduce certainty or severity accordingly.
