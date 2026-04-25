// Package bridge provides a client for the Polymarket Bridge API.
//
// # Overview
//
// The Bridge API manages cross-chain USDC bridge configuration, allowing users
// to configure their preferred bridge (Arbitrum, Base, Linea, Polygon zkEVM, etc.).
// All endpoints are public (no auth required).
//
// # Usage
//
//	client := bridge.New(bridge.Config{})  // defaults to bridge-api.polymarket.com
//
//	// Get all supported bridges
//	bridges, _ := client.GetBridges(ctx)
//
//	// Get user's configured bridge
//	config, _ := client.GetConfiguration(ctx, "0x...")
//
// Default host: https://bridge-api.polymarket.com
package bridge
