# graceful-shutdown

Status: `starter-ready`

## Purpose

Stop a long-running process cleanly on `SIGINT` or `SIGTERM`.

## Use When

- the process owns goroutines or other long-running work
- shutdown should be coordinated through `context.Context`
- you want one default shape for service termination

## Avoid When

- the program is a short-lived command with no background work
- lifecycle is already managed by a higher-level runtime

## Guarantees

- OS signals are converted into one cancellation path
- worker loops can stop without polling global state
- shutdown intent is easy to pass into downstream calls

## Tradeoffs

- every long-running operation must honor the shared context
- cancellation discipline has to be maintained across new code

## Implementation Notes

The pattern uses `signal.NotifyContext` so the same context can flow into workers and
request-scoped operations. The example keeps the worker loop small and makes shutdown
behavior obvious.

## Copy Surface

Copy the signal-derived context setup and the worker loop shape. Rename log messages and
replace `doWork` with real I/O.

## Related

- `../worker-group`
- `../http-server-lifecycle`
