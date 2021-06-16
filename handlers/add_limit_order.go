package handlers

import (
	"encoding/json"
	"github.com/conorbros/orderbook-server/orderbook"
	"github.com/shopspring/decimal"
	"net/http"
)

// AddLimitOrderRequest represents the JSON body that the limit order handler expects
type AddLimitOrderRequest struct {
	Asset    string
	Side     OrderSide
	OrderID  string
	Quantity decimal.Decimal
	Price    decimal.Decimal
}

// AddLimitOrderResponse represents the body that the limit order handler will respond with
type AddLimitOrderResponse struct {
	Done                     []*orderbook.Order
	Partial                  *orderbook.Order
	PartialQuantityProcessed decimal.Decimal
}

// AddLimitOrderHandler handles a request to add a limit order to the orderbook
func AddLimitOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Body is empty.", http.StatusBadRequest)
		return
	}

	var req AddLimitOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Malformed JSON.", http.StatusBadRequest)
		return
	}

	if Assets[req.Asset] == nil {
		AddAsset(req.Asset)
	}

	var side orderbook.Side
	if req.Side == ASK {
		side = orderbook.Sell
	} else if req.Side == BID {
		side = orderbook.Buy
	} else {
		http.Error(w, "Invalid order side.", http.StatusBadRequest)
		return
	}

	Lock.Lock()
	done, partial, partialQuantityProcessed, err := Assets[req.Asset].ProcessLimitOrder(side, req.OrderID, req.Quantity, req.Price)
	Lock.Unlock()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := AddLimitOrderResponse{
		done,
		partial,
		partialQuantityProcessed,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
