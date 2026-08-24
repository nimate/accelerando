# accelerando

Routing rules for agents working in this repo.

## Load Minimally

- Do not load all patterns or starters by default.
- Start with [`docs/architecture/catalog.md`](docs/architecture/catalog.md).
- Open only the specific pattern or starter directly relevant to the task.
- Read [`RECAP.md`](RECAP.md) only when you need current repo intent or canonical
  boundaries.

## Canonical Paths

- `patterns/go/` is the canonical source for focused reusable concerns.
- `starters/go/` is the canonical source for runnable service shapes.
- `docs/architecture/` is the canonical navigation and repo-structure surface.
- `_playground/` is non-canonical scratch/import material and should not be treated as
  the source of truth.

## How To Choose

- Use `starters/go/` when starting a new service.
- Use `patterns/go/` when implementing or reviewing one focused behavior.
- If starter code contains a reusable concern, promote that concern into `patterns/go/`
  instead of letting the starter become a hidden pattern library.

## Working Style

- Prefer local repo patterns over inventing architecture from scratch.
- Keep context usage tight: read the smallest useful set of files.
- When adding a new pattern or starter, keep its README short and retrieval-friendly.

## Local Agent Safety Policy

Agents may create files and directories, edit files, and run local tests, `make`, and
other project tooling without asking for routine confirmation.

Before deleting an existing file or directory, agents must ask the user for confirmation.
This includes `rm`, `rmdir`, `git clean`, `git reset`, `git checkout`, and `git restore`
when they would remove or overwrite existing work.
