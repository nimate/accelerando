# structured-logging

Status: `validated`

## Purpose

Establish one local default for structured application logs using `log/slog`.

## Use When

- services should log key/value records consistently
- the codebase needs a shared logger shape
- you want startup code to inject a logger instead of using globals

## Avoid When

- the service already standardizes on another logging library
- logs are temporary and local-only

## Guarantees

- logs carry stable field names
- the logger is easy to pass through application code
- human-readable and machine-readable formats can be swapped centrally

## Tradeoffs

- all packages need to accept or derive a logger
- teams must agree on field naming to avoid drift

## Implementation Notes

The pattern uses `slog.NewJSONHandler` and sets the logger once in `main`. Service code
accepts `*slog.Logger` explicitly instead of calling package-level logging helpers.

## Copy Surface

Copy `newLogger` and the dependency injection pattern into your `main` package.

## Related

- `../config-loading`
- `../../../../starters/go/worker-service`
