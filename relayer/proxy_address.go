package relayer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const proxyImplementationPolygon = "0x1f5e5d53ea693e7e60761fd565bbf1eb7fe77032"

// DeriveProxyWalletAddress derives the Polymarket proxy wallet address for a signer.
//
// This mirrors the proxy-wallet flow used by Polymarket Magic users.
func DeriveProxyWalletAddress(owner common.Address) common.Address {
	factory := common.HexToAddress(proxyFactoryPolygon)
	implementation := common.HexToAddress(proxyImplementationPolygon)

	salt := crypto.Keccak256Hash(owner.Bytes())
	initCodeHash := proxyMinimalProxyInitCodeHash(implementation)

	var buf []byte
	buf = append(buf, 0xff)
	buf = append(buf, factory.Bytes()...)
	buf = append(buf, salt.Bytes()...)
	buf = append(buf, initCodeHash.Bytes()...)

	return common.BytesToAddress(crypto.Keccak256(buf)[12:])
}

func proxyMinimalProxyInitCodeHash(implementation common.Address) common.Hash {
	// EIP-1167 minimal proxy creation code:
	// 0x3d602d80600a3d3981f3
	// runtime:
	// 0x363d3d373d3d3d363d73 <implementation> 0x5af43d82803e903d91602b57fd5bf3
	code := []byte{
		0x3d, 0x60, 0x2d, 0x80, 0x60, 0x0a, 0x3d, 0x39, 0x81, 0xf3,
		0x36, 0x3d, 0x3d, 0x37, 0x3d, 0x3d, 0x3d, 0x36, 0x3d, 0x73,
	}
	code = append(code, implementation.Bytes()...)
	code = append(code, []byte{
		0x5a, 0xf4, 0x3d, 0x82, 0x80, 0x3e, 0x90, 0x3d, 0x91, 0x60,
		0x2b, 0x57, 0xfd, 0x5b, 0xf3,
	}...)

	return crypto.Keccak256Hash(code)
}
