# Trading Orders

## Placing a Single Order

```go
resp, err := client.PostOrder(ctx, clob.PostOrderRequest{
    TokenID:       "0xtoken...",
    Price:         clob.Float64{Value: 0.65},
    Size:          clob.Float64{Value: 50.0},
    Side:          clob.SideBuy,  // or clob.SideSell
    Expiration:    "GTC",         // Good Till Cancel
    OrderType:     clob.OrderTypeLmt,
    ClientMetadata: 0,
})
```

## Batch Orders

```go
responses, err := client.PostOrders(ctx, []clob.PostOrderRequest{
    {TokenID: "...", Price: clob.Float64{Value: 0.5}, Size: clob.Float64{Value: 10}, Side: clob.SideBuy},
    {TokenID: "...", Price: clob.Float64{Value: 0.6}, Size: clob.Float64{Value: 20}, Side: clob.SideBuy},
}, false, false) // postOnly=false, deferExec=false
```

## Cancelling Orders

```go
// Single order
resp, err := client.CancelOrder(ctx, "order-id")

// Multiple orders
resp, err := client.CancelOrders(ctx, []string{"id1", "id2"})

// All orders for the user
resp, err := client.CancelAll(ctx)

// Cancel orders for a specific market
resp, err := client.CancelMarketOrders(ctx, clob.OrderMarketCancelParams{
    Market: "0xconditionID",
})
```

## Querying Orders

```go
// Specific order
order, err := client.GetOrder(ctx, "order-id")

// All open orders
orders, err := client.GetOpenOrders(ctx, clob.OpenOrderParams{
    Market: "0xconditionID",
    AssetID: "0xtokenID",
})

// User's trade history
trades, err := client.GetTrades(ctx, clob.TradeParams{
    Market:  "0xconditionID",
    Limit:   100,
})
```

## Checking Order Scoring

```go
// Single order
scoring, err := client.IsOrderScoring(ctx, "order-id")

// Multiple orders
scores, err := client.AreOrdersScoring(ctx, []string{"id1", "id2"})
```
