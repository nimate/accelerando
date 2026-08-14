# go-ethereum Adapter

Use this variant when the request names `go-ethereum` or the repository already depends on it.

## Boundary

- Depend on the smallest interface containing the `ethclient` methods currently used, such as `HeaderByNumber` or `FilterLogs`.
- Accept the narrow interface in an internal constructor for tests and return the concrete wrapper from the production constructor.
- Dial in the production constructor, prefer a context-aware dial path when startup cancellation matters, and expose `Close` for the owned connection.
- Keep `types.Header`, `types.Log`, `common.Hash`, `common.Address`, `hexutil`, and `rpc` types inside the adapter unless callers explicitly require native go-ethereum interoperability.
- Map results into stable domain types and guard against nil headers or missing required fields before dereferencing.

## Ethereum values and hex

- Prefer go-ethereum's typed APIs and protocol types over strings: `common.Address`, `common.Hash`, `common.BytesToHash`, `common.HexToAddress` only when its permissive conversion semantics are appropriate, `types.*`, and `rpc.BlockNumber`.
- Use `github.com/ethereum/go-ethereum/common/hexutil` for JSON-RPC quantity and data encoding/decoding (`EncodeUint64`, `Uint64`, `Big`, `Bytes`, and related types). Let `encoding/json` marshal/unmarshal those types where it owns the RPC wire representation.
- Do not write local hex encoders, decoders, fixed-width padding, `0x` stripping, or `strconv.FormatUint(..., 16)` helpers for Ethereum RPC values. If go-ethereum has no suitable helper, first check the repository's approved `hexutils` package; otherwise add a narrowly tested dependency or helper only with a documented semantic need.
- Distinguish RPC **quantities** (minimal `0x`-prefixed integer representation) from **data** (byte strings, where leading zero bytes are significant). Do not use address/hash conversion helpers to encode either.
- Validate externally supplied addresses, hashes, and byte strings according to the consumer contract. Do not silently turn malformed input into a zero address/hash.

## Request semantics

- Use `HeaderByNumber` when only header fields are required instead of fetching a full block.
- Use the official `rpc.BlockNumber` sentinels for named tags such as `safe` and `finalized`; verify behavior against the installed go-ethereum version.
- Validate user-provided addresses and topic hashes before making calls.
- Preserve Ethereum filter semantics for nil block bounds and topic alternatives.
- Treat subscriptions separately from request/response calls. Reconnect, resubscribe, backfill gaps, and handle removed logs or reorgs only when required by the consumer.

## Retry

Do not retry every error returned by `ethclient`. Build a classifier around error types exposed by the installed go-ethereum version plus the provider's documented JSON-RPC codes. Verify current types in module source or official documentation instead of relying on remembered names.

Read-only calls such as headers, balances, and logs are normally candidates for transient retry. Transaction submission and other state-changing workflows require an operation-specific idempotency decision.

## Tests

- Unit-test wrapper behavior with a fake implementing the narrow interface.
- Inject backoff and clock behavior so tests do not sleep and timestamps are deterministic.
- Add an in-process or `httptest` JSON-RPC test when this code relies on go-ethereum request encoding, named block tags, provider errors, or transport classification.
- Keep public-node compatibility tests opt-in.
