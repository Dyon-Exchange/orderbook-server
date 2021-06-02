package main

import (
	"io"
	"log"
	"net/http"

	"github.com/conorbros/orderbook-server/handlers"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	http.HandleFunc("/addAsset", handlers.AddAssetHandler)

	http.HandleFunc("/addLimitOrder", handlers.AddLimitOrderHandler)

	http.HandleFunc("/addMarketOrder", handlers.AddMarketOrderHandler)

	http.HandleFunc("/cancelOrder", handlers.CancelOrderHandler)

	http.HandleFunc("/calculateMarketPrice", handlers.CalculateMarketPriceHandler)

	http.HandleFunc("/healthCheck", HealthCheckHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
