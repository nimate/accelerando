# Artifact Status

Every pattern and starter should declare its maturity near the top of its `README.md`.

## Status Values

- `draft`: imported or experimental material, not yet a recommended default
- `validated`: reviewed and considered a good local default
- `starter-ready`: safe to compose into starter services
- `deprecated`: kept only for reference or migration

## Usage Rule

- A new idea can start as `draft`.
- A reusable pattern should not be composed into a starter until it is at least
  `validated`.
- Starters should only depend on patterns that are `validated` or `starter-ready`.
