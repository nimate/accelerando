# cli-tool

Status: `validated`

Runnable CLI starter for small Go commands.

## Composes

- `patterns/go/config-loading`
- `patterns/go/structured-logging`

## Includes

- argument parsing through the standard library `flag` package
- structured logging
- explicit exit path from `main`

## Intentionally Left Out

- subcommand frameworks
- config files
- background workers
- interactive TUI behavior

## Use This When

- the program is short-lived and request/response shaped
- you want the smallest starter in the repo

## Choose Something Else When

- the process owns long-running goroutines
- HTTP lifecycle is part of the product
