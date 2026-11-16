# Accelerando Catalog

## Build A New Service

- HTTP API in Go: `starters/go/api-service`
- Background worker in Go: `starters/go/worker-service`
- CLI tool in Go: `starters/go/cli-tool`

## Reuse A Focused Behavior

- Graceful shutdown: `patterns/go/graceful-shutdown`
- Cancellable worker groups: `patterns/go/worker-group`
- HTTP server lifecycle: `patterns/go/http-server-lifecycle`
- Config from environment: `patterns/go/config-loading`
- Structured logs with `log/slog`: `patterns/go/structured-logging`
- Burst-tolerant rate limiting: `patterns/go/rate-limiting`

## Composition Rules

- Prefer a starter when beginning a new service.
- Prefer a pattern when modifying an existing codebase.
- If a starter introduces reusable behavior, extract or update the pattern first.
