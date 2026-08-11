---
name: mac-review-go-rpc-client
description: Review and audit existing Go RPC client implementations and their tests. Use when asked to review, assess, validate, or find bugs in RPC, JSON-RPC, HTTP API, blockchain, Ethereum, or go-ethereum client code, with emphasis on behavioral regressions, interface boundaries, context and timeout handling, error classification, retry and backoff correctness, observability, lifecycle cleanup, concurrency, security, and missing tests. Do not use for initial client implementation unless the user explicitly asks for both review and fixes.
---

# Review Go RPC Client

Review behavior, not stylistic preference. Lead with concrete defects and risks, ordered by severity and supported by file and line references. Do not edit code unless the user asks for fixes.

## Establish scope

1. Read repository instructions and determine the requested review boundary: files, package, diff, pull request, or full client subsystem.
2. Inspect `go.mod`, public consumers, constructors, configuration, domain models, transport adapters, retry helpers, logging, tests, and call sites.
3. Identify the protocol and concrete transport from evidence. Do not ask the user to choose a library when reviewing an implementation that already made that choice.
4. Establish the intended behavior for every reviewed operation: inputs, outputs, idempotency, timeout ownership, retry policy, concurrency promise, and resource ownership.
5. Use documentation and installed dependency source when correctness depends on a version-specific SDK or protocol contract. State when a conclusion is an inference.

Do not report an issue merely because the code differs from a preferred pattern. Trace a plausible failure mode, caller-visible regression, security consequence, operational problem, or test gap.

## Inspect systematically

Read [review-checklist.md](references/review-checklist.md) and apply the relevant sections. Read [testing-review.md](references/testing-review.md) when evaluating coverage or running tests. For a `go-ethereum` adapter, also read [go-ethereum-review.md](references/go-ethereum-review.md).

Prioritize these questions:

1. Does the public contract match what callers need without accidental SDK coupling?
2. Can nil, malformed, partial, or protocol-error responses panic or produce a false success?
3. Are caller cancellation and deadlines honored during dialing, calls, and backoff waits?
4. Are retries limited to transient failures and operations safe to repeat?
5. Are maximum attempts, elapsed time, jitter, caps, and server retry hints implemented as documented?
6. Can shared state race under concurrent use?
7. Are connections, response bodies, subscriptions, timers, and goroutines cleaned up?
8. Can endpoints, credentials, headers, payloads, or customer data leak through logs and errors?
9. Do tests exercise failure semantics without real sleeps or public network dependencies?

## Validate findings

For each candidate finding:

1. Locate the exact code path and affected caller behavior.
2. Check nearby tests and repository-wide references before concluding coverage is absent.
3. Distinguish a demonstrated bug from a conditional risk or open question.
4. Reproduce with a focused test or existing command when practical and safe.
5. Avoid broad refactoring recommendations when a narrow correction addresses the risk.

Use severity consistently:

- **Critical**: likely security compromise, irreversible corruption, or widespread outage.
- **High**: plausible production failure, duplicated mutation, persistent resource leak, or broken core behavior.
- **Medium**: incorrect behavior under a realistic edge case, weakened resilience, or material observability gap.
- **Low**: limited-impact robustness issue or maintainability problem with a concrete failure path.

Do not inflate severity because a client handles network or blockchain traffic.

## Run verification

Prefer repository-prescribed commands. For Go, use the applicable subset:

```sh
go vet ./...
go test ./...
go test -race ./...
```

Add or propose a focused reproducer only when needed to validate a suspected bug. Do not use live RPC endpoints in the default verification path. Report commands that could not run and why.

## Report findings first

Present findings before summaries, ordered from highest to lowest severity. Use this shape:

```markdown
- **[High] Short imperative title** - `path/file.go:42`
  Explain the triggering condition, observable impact, and why existing handling or tests do not prevent it. Suggest the smallest viable direction for correction.
```

Keep each finding self-contained. Cite the narrowest useful line. Group multiple locations only when they form one defect.

After findings, include only the sections that add value:

1. **Open Questions** for assumptions that block a firm conclusion.
2. **Test Gaps** for important behavior that remains unverified but is not itself a demonstrated bug.
3. **Summary** as brief secondary context.

If there are no findings, say so explicitly and state remaining test gaps or residual risks. Do not manufacture findings to fill the report.
