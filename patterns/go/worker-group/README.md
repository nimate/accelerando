# worker-group

Status: `starter-ready`

## Purpose

Run a group of cancellable goroutines that share one lifecycle and fail together.

## Use When

- multiple background loops belong to one process
- one worker failure should stop the group
- you want to tie worker lifetime to `context.Context`

## Avoid When

- workers are independent and should outlive each other
- a plain sequential flow would be simpler

## Guarantees

- all workers see the same cancellation signal
- the first error cancels sibling workers
- the caller waits for the whole group to settle

## Tradeoffs

- callers must decide which errors are terminal
- cancellation can mask follow-up errors in sibling workers

## Implementation Notes

`errgroup.WithContext` is the default when a service owns a small set of long-running
workers. The example uses one ticker-driven loop and one consumer loop to show shared
shutdown semantics.

## Copy Surface

Copy the `errgroup.WithContext` setup and worker function shape. Replace the toy work with
real blocking operations that accept `ctx`.

## Related

- `../graceful-shutdown`
- `../rate-limiting`
