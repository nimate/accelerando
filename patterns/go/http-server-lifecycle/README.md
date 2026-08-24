# http-server-lifecycle

Status: `starter-ready`

## Purpose

Start an HTTP server with explicit timeouts and shut it down cleanly.

## Use When

- building an HTTP API or internal admin server
- startup and shutdown behavior should be explicit
- you want one local default for server timeouts

## Avoid When

- the runtime fully owns HTTP server lifecycle
- you are building a one-shot CLI that only makes outgoing requests

## Guarantees

- listen and serve errors are surfaced
- shutdown has a bounded timeout
- basic server hardening is encoded in one place

## Tradeoffs

- more boilerplate than `http.ListenAndServe`
- timeout values still need per-service tuning

## Implementation Notes

The pattern constructs `http.Server` directly instead of hiding behavior in helper
functions. Timeouts are explicit so starters can inherit the same shape.

## Copy Surface

Copy the `http.Server` initialization, the serve goroutine, and the shutdown timeout flow.
Replace handlers and timeout values to fit the service.

## Related

- `../graceful-shutdown`
- `../../../../starters/go/api-service`
