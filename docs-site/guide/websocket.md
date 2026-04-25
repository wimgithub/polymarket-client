# WebSocket Streams

The WebSocket client provides real-time order book and order update streams.

## Connection

```go
import "github.com/bububa/polymarket-client/clob/ws"

// Read-only (order book only)
wsClient, err := ws.New(ws.Config{Host: ""})

// With authentication (order book + order notifications)
wsClient, err := ws.New(ws.Config{
    Signer:      polyauth.NewSigner(privateKey),
    Credentials: &ws.Credentials{Key: "...", Secret: "...", Passphrase: "..."},
    ChainID:     137,
})
defer wsClient.Close()
```

## Subscriptions

```go
// Subscribe to order book snapshots for a token
err := wsClient.SubscribeOrderBook("token-id")

// Subscribe to order update notifications (requires auth)
err = wsClient.SubscribeOrders()
```

## Reading Updates

All updates arrive on `wsClient.Channel`:

```go
for update := range wsClient.Channel {
    switch update.Type {
    case "book":
        // Order book snapshot
    case "order":
        // Order fill or status change
    }
}
```

## Reconnection

The client does **not** auto-reconnect. On `ErrConnectionLost`:

```go
err := wsClient.Read()
if err == ws.ErrConnectionLost {
    wsClient.Close()
    // Create new client, authenticate, and re-subscribe
}
```
