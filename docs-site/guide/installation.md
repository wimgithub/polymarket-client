# Installation

Requires **Go 1.23+**.

```bash
go get github.com/bububa/polymarket-client
```

## Go Module

```go
import (
    "github.com/bububa/polymarket-client/clob"
    "github.com/bububa/polymarket-client/data"
    "github.com/bububa/polymarket-client/gamma"
    "github.com/bububa/polymarket-client/relayer"
    "github.com/bububa/polymarket-client/bridge"
    "github.com/bububa/polymarket-client/clob/ws"
    "github.com/bububa/polymarket-client/shared"
)
```

## Verify Installation

```bash
go build -v ./...
go test -v ./...
```

All tests use `httptest.NewServer` and require no network access.

## Go Version Note

The `go.mod` file declares `go 1.22` for backward compatibility, but the CI pipeline uses `>=1.23.0`. We recommend Go 1.23+ for the best experience.
