# Recap

Read this first if you are a fresh AI session entering the repo.

## What This Project Is For

`accelerando` is an opinionated Go-first project that helps bootstrap new services with
local patterns and starters.

The core idea is:

- do not rely on the model to rediscover good service structure from scratch
- keep repeatable implementation decisions in-repo
- make those decisions easy to retrieve for both a human and an LLM

This is a real standalone project now, not a leftover interview-prep template.

## How To Read The Repo

There are two canonical layers:

- `patterns/go/` contains one reusable concern per directory
- `starters/go/` contains runnable service shapes composed from those concerns

Read `patterns/` when the task is "how do we do X here?"
Read `starters/` when the task is "start a new service of type Y."

The main retrieval index is [`docs/architecture/catalog.md`](docs/architecture/catalog.md).

## What Counts As Canonical

Canonical:

- `patterns/go/*`
- `starters/go/*`
- `docs/architecture/*`

Non-canonical:

- `_playground/`

`_playground/` may contain useful seeds or copies from other projects, but it is not the
source of truth and should not be treated as the recommended implementation location.

## Operating Rule

If code inside a starter expresses a reusable concern, promote it into `patterns/go/`
and document it there. Starters should stay thin and compositional.

If a pattern is only an experiment, keep it out of the canonical surface until it has a
clear README and a stable intended use.

## What Good Work Looks Like Here

Good additions usually have these properties:

- one concern, one directory, one clear README
- runnable or nearly runnable Go code
- comments that explain design choices, not syntax
- obvious retrieval paths for future prompts and future humans

When in doubt, prefer making the local pattern library clearer over making the prompt
more clever.
