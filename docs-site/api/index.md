# API Reference

Quick reference for all packages.

## relayer

**Host:** `https://relayer-v2.polymarket.com` | **Auth:** API key

| Method | Description |
|---|---|
| `SubmitTransaction` | Submit signed transaction |
| `GetTransaction` | Get transaction by ID |
| `GetRecentTransactions` | Recent transactions |
| `GetNonce` | Get relayer nonce |
| `GetRelayPayload` | Get relay payload |
| `IsSafeDeployed` | Check Safe wallet status |
| `GetAPIKeys` | List relayer API keys |

## data

**Host:** `https://data-api.polymarket.com` | **Auth:** None

| Method | Description |
|---|---|
| `GetHealth` | API health check |
| `GetPositions` | User positions |
| `GetMarketPositions` | Market-level positions |
| `GetClosedPositions` | Closed positions |
| `GetValue` | Portfolio value |
| `GetTrades` | User/market trades |
| `GetActivity` | User activity |
| `GetHolders` | Top holders |
| `GetTraded` | Markets traded count |
| `GetOpenInterest` | Market open interest |
| `GetLiveVolume` | Live volume |
| `GetLeaderboard` | Trader rankings |
| `GetBuilderLeaderboard` | Builder rankings |
| `GetBuilderVolume` | Builder volume |
| `DownloadAccountingSnapshot` | ZIP snapshot |

## gamma

**Host:** `https://gamma-api.polymarket.com` | **Auth:** None

| Method | Description |
|---|---|
| `GetMarket` / `GetMarketBySlug` | Market by ID or slug |
| `GetMarkets` | Filtered markets |
| `GetEvent` / `GetEventBySlug` | Event by ID or slug |
| `GetEvents` | Filtered events |
| `Search` | Full-text search |
| `ListSeries` / `GetSeries` | Series data |
| `GetTags` / `GetTag` / `GetTagBySlug` | Tag management |
| `GetRelatedTags` / `GetRelatedTagRelationships` | Related tags |
| `GetSports` / `GetValidSportsMarketTypes` | Sports data |
| `GetTeams` | Sports teams |
| `GetComments` / `GetComment` / `GetCommentsByUserAddress` | Comments |
| `GetPublicProfile` | User profile by address |

## bridge

**Host:** `https://bridge-api.polymarket.com` | **Auth:** None

| Method | Description |
|---|---|
| `GetBridges` | List supported bridges |
| `GetConfiguration` | User's bridge config |

## shared

| Type | Purpose |
|---|---|
| `Float64` | Always unmarshals as float64 (accepts strings, numbers, null) |
| `String` | Converts strings, numbers, and booleans into a stable string form |
| `Time` / `Date` | Accept common Polymarket timestamp and date encodings |
| `StringSlice` / `Float64Slice` | Accept arrays and string-encoded arrays |
