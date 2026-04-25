# Relayer Integration

The Relayer API submits pre-signed EIP-712 transactions to Polygon on the user's behalf.

## Direct Relayer Usage

```go
import "github.com/bububa/polymarket-client/relayer"

relayerClient := relayer.New(relayer.Config{
    Credentials: &relayer.Credentials{
        APIKey:  "your-relayer-api-key",
        Address: "0xYourWalletAddress",
    },
})

// Submit a transaction
resp, err := relayerClient.SubmitTransaction(ctx, relayer.SubmitTransactionRequest{
    Type:     "order",
    Payload:  signedPayload,
    Signature: sig,
})
```

## Through CLOB Client

The CLOB client can delegate CTF transactions to a configured Relayer:

```go
clobClient := clob.NewClient("",
    clob.WithRelayerClient(relayerClient),
)

// This internally calls relayerClient.SubmitTransaction
resp, err := clobClient.SubmitRelayerTransaction(ctx, req)
```

## Querying Transactions

```go
// By ID
tx, err := relayerClient.GetTransaction(ctx, "tx-id")

// Recent transactions (requires auth)
txs, err := relayerClient.GetRecentTransactions(ctx)

// Check Safe wallet deployment
deployed, err := relayerClient.IsSafeDeployed(ctx, "0x...")
```
