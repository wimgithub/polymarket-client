package clob

import (
	"context"
	"fmt"
)

type OrderArgsV2 struct {
	TokenID       string
	Price         string
	Size          string
	Side          Side
	Expiration    string
	SignatureType SignatureType
	BuilderCode   string
	Metadata      string
}

type MarketOrderArgsV2 struct {
	TokenID       string
	Price         string // worst-price limit (optional, required for amount→shares conversion)
	Amount        string // BUY: USDC (pUSD) to spend, SELL: shares to sell
	Side          Side
	SignatureType SignatureType
	BuilderCode   string
	Metadata      string
}

type CreateOrderOptions struct {
	TickSize string // required tick size; price will be validated, not auto-rounded
	NegRisk  bool
}

type OrderBuilder struct {
	client *Client
}

func NewOrderBuilder(client *Client) *OrderBuilder {
	return &OrderBuilder{client: client}
}

func (b *OrderBuilder) BuildOrder(args OrderArgsV2, opts CreateOrderOptions) (*SignedOrder, error) {
	err := ValidateBytes32Hex("builder", args.BuilderCode)
	if err != nil {
		return nil, err
	}
	err = ValidateBytes32Hex("metadata", args.Metadata)
	if err != nil {
		return nil, err
	}
	if err := validatePriceTicks(args.Price, opts.TickSize); err != nil {
		return nil, err
	}

	maker, taker, err := computeOrderAmounts(args.Price, args.Size, args.Side)
	if err != nil {
		return nil, err
	}

	order := &SignedOrder{
		TokenID:       String(args.TokenID),
		MakerAmount:   String(maker),
		TakerAmount:   String(taker),
		Side:          args.Side,
		SignatureType: args.SignatureType,
		Builder:       args.BuilderCode,
		Metadata:      args.Metadata,
		Expiration:    String("0"),
	}

	if args.Expiration != "" {
		order.Expiration = String(args.Expiration)
	}

	if err := b.client.SignOrder(order, WithSignOrderNegRisk(opts.NegRisk)); err != nil {
		return nil, err
	}
	return order, nil
}

func (b *OrderBuilder) BuildMarketOrder(args MarketOrderArgsV2, opts CreateOrderOptions) (*SignedOrder, error) {
	err := ValidateBytes32Hex("builder", args.BuilderCode)
	if err != nil {
		return nil, err
	}
	err = ValidateBytes32Hex("metadata", args.Metadata)
	if err != nil {
		return nil, err
	}
	if err := validatePriceTicks(args.Price, opts.TickSize); err != nil {
		return nil, err
	}

	maker, taker, err := computeMarketOrderAmounts(args.Price, args.Amount, args.Side)
	if err != nil {
		return nil, err
	}

	order := &SignedOrder{
		TokenID:       String(args.TokenID),
		MakerAmount:   String(maker),
		TakerAmount:   String(taker),
		Side:          args.Side,
		SignatureType: args.SignatureType,
		Expiration:    String("0"),
		Builder:       args.BuilderCode,
		Metadata:      args.Metadata,
	}

	if err := b.client.SignOrder(order, WithSignOrderNegRisk(opts.NegRisk)); err != nil {
		return nil, err
	}
	return order, nil
}

func (b *OrderBuilder) CreateAndPostOrder(ctx context.Context, args OrderArgsV2, opts CreateOrderOptions, orderType OrderType, deferExec *bool) (*PostOrderResponse, error) {
	order, err := b.BuildOrder(args, opts)
	if err != nil {
		return nil, err
	}
	req := PostOrderRequest{
		Order:     *order,
		Owner:     order.Maker,
		OrderType: orderType,
		DeferExec: deferExec,
	}
	out := &PostOrderResponse{}
	if err := b.client.PostOrder(ctx, req, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (b *OrderBuilder) CreateAndPostMarketOrder(ctx context.Context, args MarketOrderArgsV2, opts CreateOrderOptions, orderType OrderType, deferExec *bool) (*PostOrderResponse, error) {
	if orderType != FOK && orderType != FAK {
		return nil, fmt.Errorf("polymarket: market order type must be FOK or FAK, got %s", orderType)
	}
	order, err := b.BuildMarketOrder(args, opts)
	if err != nil {
		return nil, err
	}
	req := PostOrderRequest{
		Order:     *order,
		Owner:     order.Maker,
		OrderType: orderType,
		DeferExec: deferExec,
	}
	out := &PostOrderResponse{}
	if err := b.client.PostOrder(ctx, req, out); err != nil {
		return nil, err
	}
	return out, nil
}
