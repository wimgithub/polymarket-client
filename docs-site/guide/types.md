# Working with Types

## Float64 vs String

Polymarket's API returns decimal prices as strings (`"0.50"`) in some endpoints and raw numbers (`0.5`) in others. The SDK provides two types to handle this:

### Float64

Always unmarshals to `float64`, regardless of input format:

```go
type Price struct {
    Value float64
}

// Accepts: "0.50" → 0.5, 0.5 → 0.5, null → 0
var p shared.Float64
json.Unmarshal([]byte(`"0.50"`), &p) // p.Value = 0.5
```

### String

Converts strings, numbers, and booleans into a stable string form:

```go
var id shared.String
json.Unmarshal([]byte(`12345`), &id) // id.String() == "12345"
```

Use `String` for IDs, nonces, or fields that are documented as strings but sometimes arrive as JSON numbers.

## Page[T] Pagination

Most list endpoints return paginated results:

```go
type Page[T any] struct {
    Data       []T
    NextCursor string
    Limit      shared.Int
    Count      shared.Int
}

// Paginate through all results
cursor := ""
for {
    page, err := client.GetMarkets(ctx, cursor)
    if err != nil { break }
    
    for _, market := range page.Data {
        fmt.Println(market.ConditionID)
    }
    
    if page.NextCursor == "" { break }
    cursor = page.NextCursor
}
```

## URL Query Parameters

Request structs use `url` tags for automatic query parameter encoding:

```go
type TradeParams struct {
    User  string `url:"user,omitempty"`
    Limit int    `url:"limit,omitempty"`
}
// → ?user=0x...&limit=100
```

Empty/zero values are omitted from the query string.
