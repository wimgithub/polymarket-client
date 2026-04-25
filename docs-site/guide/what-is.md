# What is Polymarket Go Client?

`polymarket-client` is a pure-Go SDK for interacting with the Polymarket prediction market platform. It wraps all five Polymarket HTTP APIs and one WebSocket API.

## Key Features

- Complete [CLOB v2](api/clob) coverage (65+ methods)
- [WebSocket](api/ws) live order book and order updates
- Three-tier authentication (public → L1 wallet → L2 full trading)
- Zero live dependencies — all tests run offline
- One external dependency: `go-ethereum`

## Which Package Should I Use?

| I want to... | Package | Auth |
|---|---|---|
| Query market data, order books, prices | `clob` | None |
| Place or cancel orders | `clob` | L2 |
| Stream live order book updates | `clob/ws` | L2 (for order events) |
| Look up positions, trades, leaderboard | `data` | None |
| Search markets, events, tags, profiles | `gamma` | None |
| Submit signed on-chain transactions | `relayer` | L1 (API key) |
| Query bridge configuration | `bridge` | None |

## Next Steps

- [Installation](installation) — Add to your Go module
- [Quick Start](quick-start) — 5 minutes to your first market query
- [Auth Levels](auth-levels) — Understand the 3-tier auth system
