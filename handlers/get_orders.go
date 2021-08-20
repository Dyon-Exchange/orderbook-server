package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/conorbros/orderbook-server/orderbook"
)

// GetOrdersResponse represents the JSON body that the get orders handler will return
type GetOrdersResponse struct {
	Assets map[string]*orderbook.OrderBook
}

// GetOrdersHandler handles a request to retreive all the orders in the book
// When there is a fully fledged storage solution
func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// extract the productIdentifier from url query
	keys, ok := r.URL.Query()["prodId"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "Url Param 'prodId' is missing.", http.StatusBadRequest)
		return
	}
	key := string(keys[0])
	// Get the entire orderbook from memory
	MemQry := GetOrdersResponse{
		Assets,
	}
	// Try and locate the asset by key
	// If the asset does not have any orders return empty response
	if MemQry.Assets[key] == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("")
		return
	}
	// Otherwise only respond with orders for that asset
	res := MemQry.Assets[key]

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
