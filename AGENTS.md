# AGENTS.md — polymarket-client

Go SDK for Polymarket (prediction market on Polygon). Library only — no binaries.

## Commands

```
go build -v ./...     # compile
go test -v ./...      # tests (all use httptest, no live API needed)
go mod tidy            # clean up go.mod
```

CI uses Go >=1.23.0. `go.mod` says 1.22 — prefer 1.23+.

## Package Map

| Package | Purpose | Auth Required |
|---|---|---|
| `clob/` | CLOB v2 API — orders, markets, positions, RFQ | Mixed (0–2) |
| `clob/ws/` | WebSocket for live order book / order updates | L2 |
| `relayer/` | Relayer API — submit signed transactions | L1 |
| `data/` | Read-only market data API | None |
| `gamma/` | Read-only gamma market-data API | None |
| `bridge/` | Bridge API | None |
| `shared/` | Shared scalar types (`String`, `Int`, `Float64`, `Time`) | — |
| `internal/polyhttp/` | HTTP client with `AuthLevel` (0=none, 1=L1, 2=L2) | — |
| `internal/polyauth/` | EIP-712 signing, L1/L2 header generation | — |

## Client Construction Pattern

All public packages use `NewClient(host, opts...)` with functional options:

```go
// Minimal (read-only):
client := clob.NewClient("")          // defaults to CLOB V2 host
client := data.NewClient("")           // defaults to data API host
client := gamma.NewClient("")          // defaults to gamma API host

// With credentials:
client := clob.NewClient("",
    clob.WithCredentials(creds),       // L2 API key + passphrase
    clob.WithSigner(signer),           // EIP-712 signer from private key
    clob.WithChainID(PolygonChainID),  // 137
)
```

## Auth Levels

- **AuthNone (0)** — public endpoints (market data, orderbook)
- **AuthL1 (1)** — EIP-712 signed L1 headers (e.g., CreateAPIKey)
- **AuthL2 (2)** — requires both `Signer` **and** `Credentials` (API key, secret, passphrase). All order/trade endpoints.

L2 auth without `Credentials` → `"api credentials are required for level 2 authenticated request"`.

## Key Dependencies

- `github.com/ethereum/go-ethereum` — only external dep. Used for EIP-712 signing, Polygon RPC, ABI encoding.
- No code generation. No migrations. No build artifacts.

## Tests

All 5 test files use `httptest.NewServer` — run offline:

```
go test -v ./...
```

Test files: `clob/auth_test.go`, `clob/client_test.go`, `clob/ctf_test.go`, `relayer/client_test.go`, `shared/flex_test.go`.

## Conventions

- JSON serialization uses custom shared custom scalar types to handle Polymarket's string-encoded decimals.
- Pagination uses `Page[T]` with `next_cursor` and `limit` fields.
- URL query parameters use struct tags `` `url:"param_name,omitempty"` `` parsed via reflection (`clob/client.go:values()`).
- Each public package has a `doc.go` with a one-line description — use these, not assumptions, for package purpose.

## Release

GoReleaser configured with `builds: skip: true`. Tagging a release creates versioned module archives only.
