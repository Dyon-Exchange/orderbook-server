package main

import (
	"io"
	"log"
	"net/http"

	"github.com/conorbros/orderbook-server/handlers"
)

// HealthCheckHandler allows other services to check that the orderbook server is alive
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

	http.HandleFunc("/getOrders", handlers.GetOrdersHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	log.Fatal(http.ListenAndServe(":5341", nil))
}
