# L1 Authentication — Wallet Signatures

L1 authentication uses EIP-712 signatures from your Polygon wallet to verify identity.

## When It's Used

- `CreateAPIKey` — generate a new API key
- `DeriveAPIKey` — recover an existing API key

## Setup

```go
import (
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/bububa/polymarket-client/clob"
    "github.com/bububa/polymarket-client/internal/polyauth"
)

privateKey, _ := crypto.HexToECDSA("your-hex-key")
client := clob.NewClient("",
    clob.WithSigner(polyauth.NewSigner(privateKey)),
    clob.WithChainID(clob.PolygonChainID), // 137
)
```

## Creating an API Key

```go
nonce := time.Now().UnixNano() // or any incrementing int64
creds, err := client.CreateAPIKey(ctx, nonce)
// creds.Key, creds.Secret, creds.Passphrase
```

The returned `Credentials` are used for L2 requests (trading).
