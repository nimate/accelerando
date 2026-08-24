# rate-limiting

Status: `validated`

## Purpose

Allow short bursts while enforcing a steady processing rate.

## Use When

- a worker or API consumer must smooth uneven traffic
- the system should tolerate a small burst without immediate blocking
- a full external limiter library would be overkill

## Avoid When

- limits must be coordinated across multiple processes
- fairness between tenants matters more than simplicity

## Guarantees

- the burst size is explicit
- token refill rate is explicit
- a full buffer never blocks the refill loop

## Tradeoffs

- this is process-local only
- operational tuning is manual

## Implementation Notes

The pattern uses a buffered channel as a token bucket. It is intentionally small and easy
to reason about.

## Copy Surface

Copy the limiter construction and token refill goroutine. Replace the request simulation
with real inbound work.

## Related

- `../worker-group`
- `../../../../_playground/burst-limiter`
