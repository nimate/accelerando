# go-ethereum Review

Apply this checklist when the client wraps `github.com/ethereum/go-ethereum`.

## Adapter boundary

- Check that the wrapper depends on only the `ethclient` methods it uses and can substitute them in unit tests.
- Verify production dialing owns and exposes cleanup for the underlying client.
- Check context-aware dialing when startup cancellation or timeout is required.
- Confirm `types.Header`, `types.Log`, `common`, `hexutil`, and `rpc` types do not leak across a boundary intended to remain library-independent.
- Guard nil headers and required fields before dereferencing or converting them.

## Ethereum semantics

- Prefer header queries when transaction bodies are unnecessary.
- Verify named block tags such as `safe` and `finalized` use behavior supported by the installed go-ethereum version and target node.
- Validate addresses and topic hashes without silently normalizing invalid user input.
- Preserve filter semantics for nil bounds, named blocks, and topic alternatives.
- Treat removed logs, chain reorganization, subscription reconnect, and gap backfill as requirements only when the client exposes streaming or indexing behavior.

## Errors and mutations

- Do not accept blanket retry of all `ethclient` errors.
- Inspect the installed module version for actual error types and transport behavior before recommending `errors.As` checks.
- Combine go-ethereum error information with documented provider RPC codes when classifying transient failures.
- Review transaction submission and other mutations separately; a missing response does not prove the node rejected the operation.

## Tests

- Use a fake narrow interface for wrapper logic and deterministic retries.
- Use an in-process or HTTP JSON-RPC server when named-tag encoding, provider errors, or transport behavior is owned by the code under review.
- Keep public-node tests outside the default suite.
