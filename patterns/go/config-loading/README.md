# config-loading

Status: `validated`

## Purpose

Load process configuration from environment variables into a typed struct.

## Use When

- the service has a small number of required settings
- you want config validation near process startup
- env vars are the deployment contract

## Avoid When

- configuration comes from a richer source such as files or remote config
- there is no real validation to perform

## Guarantees

- required settings fail fast
- defaults are explicit in code
- the rest of the process works with typed config, not raw env lookups

## Tradeoffs

- manual parsing is repetitive for larger configs
- nested or dynamic config shapes need a different approach

## Implementation Notes

The pattern keeps parsing in one place and returns a typed `Config` plus a single startup
error. It uses the standard library only.

## Copy Surface

Copy the `Config` type and `loadConfig` function, then add or remove fields for the
service.

## Related

- `../structured-logging`
- `../../../../starters/go/api-service`
