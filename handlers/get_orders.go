package handlers

import (
	"encoding/json"

	"github.com/conorbros/orderbook-server/orderbook"
	"net/http"
)

// GetOrdersResponse represents the JSON body that the get orders handler will return
type GetOrdersResponse struct {
	Assets map[string]*orderbook.OrderBook
}

// GetOrdersHandler handles a request to retreive all the orders in the book
func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	res := GetOrdersResponse{
		Assets,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
