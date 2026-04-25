# API Key Lifecycle

## Creating

Use L1 auth to create a new API key:

```go
creds, err := client.CreateAPIKey(ctx, nonce)
```

Store `Key`, `Secret`, and `Passphrase` securely — the secret is only returned once.

## Deriving (Recovery)

If you have the original nonce, you can derive an existing key:

```go
creds, err := client.DeriveAPIKey(ctx, nonce)
```

## Listing All Keys

Requires L2 auth:

```go
keys, err := client.GetAPIKeys(ctx)
```

## Deleting

```go
err := client.DeleteAPIKey(ctx)
```

## API Key Rotation

```
1. Create new API key (L1)
2. Update your application with new credentials
3. Delete old API key (L2, using new key)
```

:::warning
Never commit API secrets to version control. Use environment variables or a secrets manager.
:::
