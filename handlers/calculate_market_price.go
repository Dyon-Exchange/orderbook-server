package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/conorbros/orderbook-server/orderbook"
	"github.com/shopspring/decimal"
)

type CalculateMarketPriceRequest struct {
	Asset    string
	Side     OrderSide
	Quantity decimal.Decimal
}

type CalculateMarketPriceResponse struct {
	Price decimal.Decimal
}

func CalculateMarketPriceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Body is empty.", http.StatusBadRequest)
		return
	}

	var req CalculateMarketPriceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Malformed JSON.", http.StatusBadRequest)
		return
	}

	if Assets[req.Asset] == nil {
		http.Error(w, "Asset does not exist.", http.StatusBadRequest)
		return
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

	price, err := Assets[req.Asset].CalculateMarketPrice(side, req.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := CalculateMarketPriceResponse{
		price,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
