# clob/ws Package

WebSocket client for live order book and order updates.

**Default Host:** `wss://ws-orderbook.clob.polymarket.com`  
**Auth:** None (order book) / L2 (order notifications)

## Creating a Client

```go
// Read-only
wsClient, err := ws.New(ws.Config{Host: ""})

// With authentication
wsClient, err := ws.New(ws.Config{
    Signer:      polyauth.NewSigner(pk),
    Credentials: &ws.Credentials{Key: "...", Secret: "...", Passphrase: "..."},
    ChainID:     137,
})
```

## Subscriptions

| Method | Auth | Description |
|---|---|---|
| `SubscribeOrderBook(tokenID)` | None | Subscribe to order book snapshots |
| `UnsubscribeOrderBook(tokenID)` | None | Unsubscribe from order book |
| `SubscribeOrders()` | L2 | Subscribe to order fill/status updates |
| `UnsubscribeOrders()` | L2 | Unsubscribe from order updates |

## Reading Updates

Updates arrive on `wsClient.Channel`:

```go
for update := range wsClient.Channel {
    fmt.Printf("Type: %s\n", update.Type) // "book" or "order"
}
```

## Error Handling

`ws.ErrConnectionLost` — connection dropped. Close and reconnect manually.
