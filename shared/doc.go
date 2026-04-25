// Package shared contains shared JSON scalar helpers used by Polymarket API packages.
//
// # Custom Types
//
// Polymarket returns numeric values as both JSON numbers and decimal strings
// (e.g. "0.50" vs 0.50). These types handle the ambiguity:
//
//   - String — preserves a stable string form for strings, numbers, and booleans
//   - Int, Int64, Uint64, and Float64 — accept quoted or native JSON numbers
//   - Time and Date — accept common Polymarket timestamp and date encodings
//   - StringSlice and Float64Slice — accept arrays and string-encoded arrays
//
// # Usage
//
//	type Market struct {
//	    ID      shared.String  `json:"id"`
//	    Volume  shared.Float64 `json:"volume"`
//	    EndDate shared.Time    `json:"endDate"`
//	}
package shared
