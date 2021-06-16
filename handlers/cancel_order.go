package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/conorbros/orderbook-server/orderbook"
	"net/http"
)

// CancelOrderRequest represents the JSON field that the handler expects to recieve
type CancelOrderRequest struct {
	OrderID string
	Asset   string
}

// CancelOrderResponse represents the JSON body that the handler will return
type CancelOrderResponse struct {
	Order *orderbook.Order
}

// CancelOrderHandler handles a request to cancel an order
func CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Body is empty.", http.StatusBadRequest)
		return
	}

	var req CancelOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Malformed JSON.", http.StatusBadRequest)
		return
	}

	if Assets[req.Asset] == nil {
		http.Error(w, "Asset does not exist.", http.StatusBadRequest)
		return
	}

	order := Assets[req.Asset].CancelOrder(req.OrderID)
	if order == nil {
		fmt.Println("Order with that id doesn't exist")
		http.Error(w, "Order with that id doesn't exist", http.StatusBadRequest)
		return
	}
	response := CancelOrderResponse{
		order,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
