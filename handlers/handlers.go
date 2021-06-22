package handlers

import (
	"github.com/conorbros/orderbook-server/orderbook"
	"sync"
)

// Assets holds a map of asset->orderbook
var Assets map[string]*orderbook.OrderBook

var Lock sync.Mutex

// OrderType represents the 2 types of orders in the orderbook
type OrderType string

const (
	// LIMIT represents a limit order
	LIMIT OrderType = "LIMIT"
	// MARKET represents a market order
	MARKET OrderType = "MARKET"
)

// OrderSide represents the 2 types of order sides in the orderbook
type OrderSide string

const (
	// ASK represents an ASK order
	ASK OrderSide = "ASK"
	// BID represents a BID order
	BID OrderSide = "BID"
)

// AddAsset adds an asset and new orderbook if doesn't exist
func AddAsset(asset string) {
	Lock.Lock()
	if Assets[asset] == nil {
		Assets[asset] = orderbook.NewOrderBook()
	}
	Lock.Unlock()
}

func init() {
	Assets = make(map[string]*orderbook.OrderBook)
}

func Reset() {
	Assets = make(map[string]*orderbook.OrderBook)
}
