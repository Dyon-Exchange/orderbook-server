package handlers

import (
	"github.com/conorbros/orderbook-server/orderbook"
)

var Assets map[string]*orderbook.OrderBook

type OrderType string

const (
	LIMIT  OrderType = "LIMIT"
	MARKET OrderType = "MARKET"
)

type OrderSide string

const (
	ASK OrderSide = "ASK"
	BID OrderSide = "BID"
)

func AddAsset(asset string) {
	if Assets[asset] == nil {
		Assets[asset] = orderbook.NewOrderBook()
	}
}

func init() {
	Assets = make(map[string]*orderbook.OrderBook)
}
