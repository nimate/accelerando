# accelerando

Accelerando is a Go-first pattern library and starter repo for AI-assisted software
development.

Its job is simple:

- keep a local catalog of good implementation patterns for recurring service concerns
- provide runnable service starters built from those patterns
- give both humans and LLMs a concrete source of truth for how new Go services should be
  shaped

The point is not to make the model invent architecture from scratch every time. The point
is to encode repeatable decisions once, keep them local, and reuse them consistently.

## Project Shape

The repo has two canonical layers:

- `patterns/go/` for narrow, reusable implementation decisions such as graceful
  shutdown, worker orchestration, config loading, structured logging, HTTP server
  lifecycle, and rate limiting
- `starters/go/` for runnable service shapes such as an HTTP API, background worker, and
  CLI, each assembled from those patterns

Supporting docs live under `docs/`, and [`docs/architecture/catalog.md`](docs/architecture/catalog.md)
is the main navigation index.

## Start Here

- On [GitHub](https://github.com/nimate/accelerando), select **Use this template** →
  **Create a new repository**. Name your project, then clone the new repository to
  work on it locally. It starts with the full Accelerando structure and its own Git
  history and remote.
- Read [`docs/architecture/catalog.md`](docs/architecture/catalog.md) for the repo map.
- Use a directory under `starters/go/` as the starting shape for your service.
- Use a directory under `patterns/go/` when you need one focused behavior in an existing
  codebase.
- Treat [`_playground/`](./_playground) as non-canonical scratch/import material.

## Pattern Contract

Each pattern directory should stay small and opinionated:

- one concern per directory
- one short `README.md` with `Purpose`, `Use When`, `Avoid When`, `Guarantees`,
  `Tradeoffs`, `Implementation Notes`, `Copy Surface`, and `Related`
- code that is runnable or very close to runnable
- comments that explain decisions, not language syntax

## Starter Contract

Each starter directory should be a thin composition of patterns:

- runnable as its own Go module
- explicit about which patterns it uses
- clear about what was intentionally left out
- easy to extend without inventing a new local architecture

## Repo Rule

If a starter grows a useful, repeatable behavior, that behavior should be promoted into
`patterns/go/` and documented there. Starters are assemblies, not a second hidden pattern
library.

## Validation

Run checks from the repository root:

```sh
make test
make verify
```

`make test` runs Python tests and Go tests across all discovered Go modules.
`make verify` runs Python tests plus changed-file Go verification and formatting.
