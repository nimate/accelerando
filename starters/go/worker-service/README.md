# worker-service

Status: `starter-ready`

Runnable background worker starter for small Go services.

## Composes

- `patterns/go/config-loading`
- `patterns/go/structured-logging`
- `patterns/go/graceful-shutdown`
- `patterns/go/worker-group`
- `patterns/go/rate-limiting`

## Includes

- one ticker-driven poller
- one queue-style worker
- shared cancellation through `errgroup.WithContext`
- structured JSON logging
- process-local burst limiter for outbound work

## Intentionally Left Out

- durable queue integrations
- retry policy abstractions
- metrics and tracing
- dependency injection frameworks

## Use This When

- the process is primarily background work
- one worker failure should stop the whole service

## Choose Something Else When

- the main surface is HTTP
- workers must be independently supervised
