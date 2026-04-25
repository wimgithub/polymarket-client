# L2 Authentication — Full Trading

L2 is required for all order-related endpoints: placing orders, canceling, checking positions, and RFQ.

## Setup

L2 requires **both** a wallet signer AND API credentials:

```go
client := clob.NewClient("",
    clob.WithCredentials(clob.Credentials{
        Key:        "api-key-from-create-api-key",
        Secret:     "api-secret-from-create-api-key",
        Passphrase: "api-passphrase-from-create-api-key",
    }),
    clob.WithSigner(polyauth.NewSigner(privateKey)),
    clob.WithChainID(clob.PolygonChainID),
)
```

## What Gets Signed

For each L2 request, the client generates:

1. **HMAC signature** — computed from the API secret, timestamp, HTTP method, URL path, and request body
2. **Wallet signature** — EIP-712 signed header bundle

## Error Messages

| Error | Cause | Fix |
|---|---|---|
| `"signer is required for authenticated request"` | `WithSigner` not set | Add `clob.WithSigner(...)` |
| `"api credentials are required for level 2 authenticated request"` | `WithCredentials` not set | Add `clob.WithCredentials(...)` |
| `"api secret decode error"` | Invalid base64 secret | Verify the secret from `CreateAPIKey` response |
