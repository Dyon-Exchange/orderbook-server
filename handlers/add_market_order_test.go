package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
)

func TestAddMarketOrder(t *testing.T) {
	body := AddMarketOrderRequest{
		Asset:    "asset21",
		Side:     "ASK",
		OrderId:  "orderid123",
		Quantity: decimal.New(2, 0),
		Price:    decimal.New(10, 0),
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/addMarketOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddMarketOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body = AddMarketOrderRequest{
		Asset:    "asset123",
		Side:     "BID",
		OrderId:  "orderid8",
		Quantity: decimal.New(2, 0),
		Price:    decimal.New(10, 0),
	}
	payload = new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err = http.NewRequest("POST", "/addMarketOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(AddMarketOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAddMarketOrderHandlerWrongMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/addMarketOrder", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddMarketOrderHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestAddMarketOrderHandlerEmptyBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/addMarketOrder", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddMarketOrderHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "Body is empty.\n" {
		t.Fatal("Invalid error message.")
	}
}

func TestAddMarketOrderHandlerMalformedJSON(t *testing.T) {
	req, err := http.NewRequest("POST", "/addMarketOrder", bytes.NewBufferString("abcdefgh"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddMarketOrderHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "Malformed JSON.\n" {
		t.Fatal("Incorrect error message")
	}
}

func TestAddMarketOrderInvalidOrderSide(t *testing.T) {
	body := AddMarketOrderRequest{
		Asset:    "asset123",
		Side:     "INVALID",
		OrderId:  "order123",
		Quantity: decimal.New(2, 0),
		Price:    decimal.New(10, 0),
	}

	payload := new(bytes.Buffer)

	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/addMarketOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddMarketOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if rr.Body.String() != "Invalid order side.\n" {
		t.Fatal("Invalid error message")
	}
}

func TestAddMarketOrderNoQuantity(t *testing.T) {
	body := AddMarketOrderRequest{
		Asset:    "asset90",
		Side:     "ASK",
		OrderId:  "orderid90",
		Quantity: decimal.New(0, 0),
		Price:    decimal.New(10, 0),
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/addMarketOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddMarketOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "orderbook: invalid order quantity\n" {
		t.Fatal("Invalid error message")
	}
}
