# RFQ (Request for Quote)

RFQ allows institutional traders to request liquidity for large orders off the public order book.

## Creating an RFQ Request

```go
req, err := client.CreateRFQRequest(ctx, clob.CreateRFQRequest{
    TokenID: "0xtoken...",
    Side:    clob.SideBuy,
    Size:    clob.Float64{Value: 10000.0},
})
```

## Listing RFQs

```go
requests, err := client.GetRFQRequests(ctx, clob.RFQListParams{
    Limit: 50,
})
```

## Creating and Accepting Quotes

```go
// Quoter creates a quote
quote, err := client.CreateRFQQuote(ctx, clob.CreateRFQQuoteRequest{...})

// Requester accepts a quote
err := client.ApproveRFQQuote(ctx, "quote-id")

// Or requester accepts a request directly
err := client.AcceptRFQRequest(ctx, "request-id")
```

## Canceling

```go
// Cancel RFQ request
err := client.CancelRFQRequest(ctx, "request-id")

// Cancel RFQ quote
err := client.CancelRFQQuote(ctx, "quote-id")
```

:::info
RFQ endpoints require L2 authentication. Both requesters and quoters must be authenticated.
:::
