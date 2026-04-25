// Package ws provides a WebSocket client for live Polymarket order book and order updates.
//
// # Connecting
//
// Basic (read-only order book):
//
//	wsClient, err := ws.New(ws.Config{Host: ""})
//
// With authentication for order notifications:
//
//	wsClient, err := ws.New(ws.Config{
//	    Signer:      polyauth.NewSigner(privateKey),
//	    Credentials: &ws.Credentials{Key: "...", Secret: "...", Passphrase: "..."},
//	    ChainID:     137,
//	})
//
// # Subscriptions
//
//	wsClient.SubscribeOrderBook("token-id")  // order book snapshots
//	wsClient.SubscribeOrders()               // order fill/status (requires auth)
//
// # Reading Updates
//
//	for update := range wsClient.Channel { ... }
//
// # Disconnection
//
// The client does NOT auto-reconnect. Handle ErrConnectionLost and re-subscribe.
package ws
