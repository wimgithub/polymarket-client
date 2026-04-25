# Troubleshooting

## Authentication Errors

### "signer is required for authenticated request"

You're trying to access an L1 or L2 endpoint without setting a signer.

```go
client := clob.NewClient("",
    clob.WithSigner(polyauth.NewSigner(privateKey)),
)
```

### "api credentials are required for level 2 authenticated request"

L2 endpoints need both signer AND credentials:

```go
client := clob.NewClient("",
    clob.WithCredentials(clob.Credentials{Key: "...", Secret: "...", Passphrase: "..."}),
    clob.WithSigner(polyauth.NewSigner(privateKey)),
)
```

## Network Issues

### Connection refused

Verify the host URL. The CLOB client defaults to `https://clob.polymarket.com` if host is empty string.

### Timeout errors

Set a custom HTTP client with longer timeout:

```go
client := clob.NewClient("",
    clob.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
)
```

## JSON Parsing

### "cannot parse server time from map[string]any"

The `/time` endpoint response format changed. Update to the latest version.

### Float values are 0

Polymarket returns prices as decimal strings (`"0.50"`). Use `shared.Float64` or `shared.String` instead of raw `float64`.

## WebSocket

### "connection lost" — no auto-reconnect

This is intentional. On `ErrConnectionLost`, close the client, create a new one, and re-subscribe.

### No updates after subscription

Call `wsClient.Read()` in a goroutine after subscribing:

```go
go func() {
    err := wsClient.Read()
    // handle error/reconnect
}()
```

## Build Issues

### "go.mod requires 1.22 but I'm using 1.21"

The CI uses Go 1.23+. Upgrade your Go toolchain.

### Import not found

Run `go mod tidy` and ensure the module path is `github.com/bububa/polymarket-client`.
