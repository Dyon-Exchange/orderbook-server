package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/conorbros/orderbook-server/orderbook"
	"github.com/shopspring/decimal"
)

func TestCalculateMarketPriceHandler(t *testing.T) {
	AddAsset("string123")
	AddAsset("string1234")
	Assets["string123"].ProcessLimitOrder(orderbook.Buy, "1", decimal.New(1000, 0), decimal.New(10, 0))

	body := CalculateMarketPriceRequest{
		Asset:    "string123",
		Side:     "ASK",
		Quantity: decimal.New(115, 0),
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/calculateMarketPrice", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculateMarketPriceHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	Assets["string1234"].ProcessLimitOrder(orderbook.Sell, "2", decimal.New(1000, 0), decimal.New(10, 0))
	body = CalculateMarketPriceRequest{
		Asset:    "string1234",
		Side:     "BID",
		Quantity: decimal.New(115, 0),
	}
	payload = new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err = http.NewRequest("POST", "/calculateMarketPrice", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(CalculateMarketPriceHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCalculateMarketPriceHandlerNoQuantity(t *testing.T) {
	AddAsset("newasset")
	body := CalculateMarketPriceRequest{
		Asset:    "newasset",
		Side:     "ASK",
		Quantity: decimal.New(115, 0),
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/calculateMarketPrice", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculateMarketPriceHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "orderbook: insufficient quantity to calculate price\n" {
		t.Errorf("Incorrect error message")
	}
}

func TestCalculateMarketPriceHandlerWrongMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/calculateMarketPrice", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculateMarketPriceHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
	if rr.Body.String() != "Method not allowed.\n" {
		t.Errorf("Incorrect error message")
	}
}

func TestCalculateMarketPriceHandlerEmptyBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/calculateMarketPrice", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculateMarketPriceHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "Body is empty.\n" {
		t.Errorf("Incorrect error message")
	}
}

func TestCalculateMarketPriceHandlerNoAsset(t *testing.T) {
	body := CancelOrderRequest{
		Asset:   "string43211",
		OrderID: "1",
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/calculateMarketPrice", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculateMarketPriceHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "Asset does not exist.\n" {
		t.Errorf("Incorrect error message")
	}
}

func TestCalculateMarketPriceHandlerGarbageJson(t *testing.T) {
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode("abc")

	req, err := http.NewRequest("POST", "/calculateMarketPrice", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculateMarketPriceHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Got wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	if rr.Body.String() != "Malformed JSON.\n" {
		t.Errorf("Incorrect error message")
	}
}

func TestMarketPriceHandlerInvalidOrderSide(t *testing.T) {
	AddAsset("asset123")
	body := CalculateMarketPriceRequest{
		Asset:    "asset123",
		Side:     "INVALID",
		Quantity: decimal.New(2, 0),
	}

	payload := new(bytes.Buffer)

	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/calculateMarketPrice", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculateMarketPriceHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "Invalid order side.\n" {
		t.Fatal("Invalid error message")
	}
}
