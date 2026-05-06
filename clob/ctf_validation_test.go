package clob

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestBuildSplitPositionTxRejectsInvalidInputs(t *testing.T) {
	client := NewClient("", WithChainID(PolygonChainID))

	tests := []struct {
		name string
		req  *SplitPositionRequest
		out  *CTFTransaction
		want string
	}{
		{
			name: "nil request",
			req:  nil,
			out:  &CTFTransaction{},
			want: "nil split position request",
		},
		{
			name: "nil output",
			req:  validSplitPositionRequest(),
			out:  nil,
			want: "nil CTF transaction output",
		},
		{
			name: "missing collateral",
			req: func() *SplitPositionRequest {
				req := validSplitPositionRequest()
				req.CollateralToken = common.Address{}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "collateral token is required",
		},
		{
			name: "missing condition id",
			req: func() *SplitPositionRequest {
				req := validSplitPositionRequest()
				req.ConditionID = common.Hash{}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "condition id is required",
		},
		{
			name: "missing partition",
			req: func() *SplitPositionRequest {
				req := validSplitPositionRequest()
				req.Partition = nil
				return req
			}(),
			out:  &CTFTransaction{},
			want: "partition is required",
		},
		{
			name: "nil partition item",
			req: func() *SplitPositionRequest {
				req := validSplitPositionRequest()
				req.Partition = []*big.Int{big.NewInt(1), nil}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "partition[1] is required",
		},
		{
			name: "zero partition item",
			req: func() *SplitPositionRequest {
				req := validSplitPositionRequest()
				req.Partition = []*big.Int{big.NewInt(1), big.NewInt(0)}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "partition[1] must be positive",
		},
		{
			name: "missing amount",
			req: func() *SplitPositionRequest {
				req := validSplitPositionRequest()
				req.Amount = nil
				return req
			}(),
			out:  &CTFTransaction{},
			want: "amount is required",
		},
		{
			name: "zero amount",
			req: func() *SplitPositionRequest {
				req := validSplitPositionRequest()
				req.Amount = big.NewInt(0)
				return req
			}(),
			out:  &CTFTransaction{},
			want: "amount must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := client.BuildSplitPositionTx(tt.req, tt.out)
			assertErrorContains(t, err, tt.want)
		})
	}
}

func TestBuildMergePositionsTxRejectsInvalidInputs(t *testing.T) {
	client := NewClient("", WithChainID(PolygonChainID))

	tests := []struct {
		name string
		req  *MergePositionsRequest
		out  *CTFTransaction
		want string
	}{
		{
			name: "nil request",
			req:  nil,
			out:  &CTFTransaction{},
			want: "nil merge position request",
		},
		{
			name: "nil output",
			req:  validMergePositionsRequest(),
			out:  nil,
			want: "nil CTF transaction output",
		},
		{
			name: "missing collateral",
			req: func() *MergePositionsRequest {
				req := validMergePositionsRequest()
				req.CollateralToken = common.Address{}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "collateral token is required",
		},
		{
			name: "missing condition id",
			req: func() *MergePositionsRequest {
				req := validMergePositionsRequest()
				req.ConditionID = common.Hash{}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "condition id is required",
		},
		{
			name: "missing partition",
			req: func() *MergePositionsRequest {
				req := validMergePositionsRequest()
				req.Partition = nil
				return req
			}(),
			out:  &CTFTransaction{},
			want: "partition is required",
		},
		{
			name: "missing amount",
			req: func() *MergePositionsRequest {
				req := validMergePositionsRequest()
				req.Amount = nil
				return req
			}(),
			out:  &CTFTransaction{},
			want: "amount is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := client.BuildMergePositionsTx(tt.req, tt.out)
			assertErrorContains(t, err, tt.want)
		})
	}
}

func TestBuildRedeemPositionsTxRejectsInvalidInputs(t *testing.T) {
	client := NewClient("", WithChainID(PolygonChainID))

	tests := []struct {
		name string
		req  *RedeemPositionsRequest
		out  *CTFTransaction
		want string
	}{
		{
			name: "nil request",
			req:  nil,
			out:  &CTFTransaction{},
			want: "nil redeem position request",
		},
		{
			name: "nil output",
			req:  validRedeemPositionsRequest(),
			out:  nil,
			want: "nil CTF transaction output",
		},
		{
			name: "missing collateral",
			req: func() *RedeemPositionsRequest {
				req := validRedeemPositionsRequest()
				req.CollateralToken = common.Address{}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "collateral token is required",
		},
		{
			name: "missing condition id",
			req: func() *RedeemPositionsRequest {
				req := validRedeemPositionsRequest()
				req.ConditionID = common.Hash{}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "condition id is required",
		},
		{
			name: "missing index sets",
			req: func() *RedeemPositionsRequest {
				req := validRedeemPositionsRequest()
				req.IndexSets = nil
				return req
			}(),
			out:  &CTFTransaction{},
			want: "index sets is required",
		},
		{
			name: "nil index set",
			req: func() *RedeemPositionsRequest {
				req := validRedeemPositionsRequest()
				req.IndexSets = []*big.Int{big.NewInt(1), nil}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "index sets[1] is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := client.BuildRedeemPositionsTx(tt.req, tt.out)
			assertErrorContains(t, err, tt.want)
		})
	}
}

func TestBuildRedeemNegRiskTxRejectsInvalidInputs(t *testing.T) {
	client := NewClient("", WithChainID(PolygonChainID))

	tests := []struct {
		name string
		req  *RedeemNegRiskRequest
		out  *CTFTransaction
		want string
	}{
		{
			name: "nil request",
			req:  nil,
			out:  &CTFTransaction{},
			want: "nil neg-risk redeem request",
		},
		{
			name: "nil output",
			req:  validRedeemNegRiskRequest(),
			out:  nil,
			want: "nil CTF transaction output",
		},
		{
			name: "missing condition id",
			req: func() *RedeemNegRiskRequest {
				req := validRedeemNegRiskRequest()
				req.ConditionID = common.Hash{}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "condition id is required",
		},
		{
			name: "missing amounts",
			req: func() *RedeemNegRiskRequest {
				req := validRedeemNegRiskRequest()
				req.Amounts = nil
				return req
			}(),
			out:  &CTFTransaction{},
			want: "amounts is required",
		},
		{
			name: "nil amount",
			req: func() *RedeemNegRiskRequest {
				req := validRedeemNegRiskRequest()
				req.Amounts = []*big.Int{big.NewInt(1), nil}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "amounts[1] is required",
		},
		{
			name: "zero amount",
			req: func() *RedeemNegRiskRequest {
				req := validRedeemNegRiskRequest()
				req.Amounts = []*big.Int{big.NewInt(1), big.NewInt(0)}
				return req
			}(),
			out:  &CTFTransaction{},
			want: "amounts[1] must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := client.BuildRedeemNegRiskTx(tt.req, tt.out)
			assertErrorContains(t, err, tt.want)
		})
	}
}

func validSplitPositionRequest() *SplitPositionRequest {
	return &SplitPositionRequest{
		CollateralToken:    ctfSafetyCollateral,
		ParentCollectionID: common.Hash{},
		ConditionID:        ctfSafetyConditionID,
		Partition:          BinaryPartition(),
		Amount:             big.NewInt(1_000_000),
	}
}

func validMergePositionsRequest() *MergePositionsRequest {
	return &MergePositionsRequest{
		CollateralToken:    ctfSafetyCollateral,
		ParentCollectionID: common.Hash{},
		ConditionID:        ctfSafetyConditionID,
		Partition:          BinaryPartition(),
		Amount:             big.NewInt(1_000_000),
	}
}

func validRedeemPositionsRequest() *RedeemPositionsRequest {
	return &RedeemPositionsRequest{
		CollateralToken:    ctfSafetyCollateral,
		ParentCollectionID: common.Hash{},
		ConditionID:        ctfSafetyConditionID,
		IndexSets:          BinaryPartition(),
	}
}

func validRedeemNegRiskRequest() *RedeemNegRiskRequest {
	return &RedeemNegRiskRequest{
		ConditionID: ctfSafetyConditionID,
		Amounts: []*big.Int{
			big.NewInt(1_000_000),
			big.NewInt(2_000_000),
		},
	}
}

func assertErrorContains(t *testing.T, err error, want string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error containing %q", want)
	}
	if !strings.Contains(err.Error(), want) {
		t.Fatalf("error = %q, want substring %q", err.Error(), want)
	}
}
