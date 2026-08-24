# api-service

Status: `starter-ready`

Runnable HTTP API starter for small Go services.

## Composes

- `patterns/go/config-loading`
- `patterns/go/structured-logging`
- `patterns/go/http-server-lifecycle`
- `patterns/go/graceful-shutdown`

## Includes

- `/healthz` endpoint
- explicit HTTP server timeouts
- environment-based config loading
- bounded shutdown timeout
- structured JSON logging

## Intentionally Left Out

- router frameworks
- metrics and tracing
- persistence and external clients
- authentication and authorization

## Use This When

- you want a plain `net/http` service as the starting point
- you need a starter that is easy for both humans and LLMs to extend

## Choose Something Else When

- a worker or CLI is the real process shape
- the service already depends on an existing internal platform framework
