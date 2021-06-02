package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
)

func TestAddLimitOrder(t *testing.T) {
	body := AddLimitOrderRequest{
		Asset:    "asset123",
		Side:     "ASK",
		OrderId:  "orderid123",
		Quantity: decimal.New(2, 0),
		Price:    decimal.New(10, 0),
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/addLimitOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddLimitOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body = AddLimitOrderRequest{
		Asset:    "asset123",
		Side:     "BID",
		OrderId:  "orderid8",
		Quantity: decimal.New(2, 0),
		Price:    decimal.New(10, 0),
	}
	payload = new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err = http.NewRequest("POST", "/addLimitOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(AddLimitOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAddLimitOrderNoQuantity(t *testing.T) {
	body := AddLimitOrderRequest{
		Asset:    "asset321",
		Side:     "ASK",
		OrderId:  "orderid123",
		Quantity: decimal.New(0, 0),
		Price:    decimal.New(10, 0),
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/addLimitOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddLimitOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if rr.Body.String() != "orderbook: invalid order quantity\n" {
		t.Fatal("Invalid error message")
	}
}

func TestAddLimitOrderInvalidOrderSide(t *testing.T) {
	body := AddLimitOrderRequest{
		Asset:    "asset123",
		Side:     "INVALID",
		OrderId:  "order123",
		Quantity: decimal.New(2, 0),
		Price:    decimal.New(10, 0),
	}

	payload := new(bytes.Buffer)

	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/addLimitOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddLimitOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if rr.Body.String() != "Invalid order side.\n" {
		t.Fatal("Invalid error message")
	}
}

func TestAddLimitOrderHandlerWrongMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/addLimitOrder", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddLimitOrderHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestAddLimitOrderHandlerEmptyBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/addLimitOrder", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddLimitOrderHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "Body is empty.\n" {
		t.Fatal("Invalid error message")
	}
}

func TestAddLimitOrderHandlerGarbageJson(t *testing.T) {
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode("abc")

	req, err := http.NewRequest("POST", "/addLimitOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddLimitOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Got wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "Malformed JSON.\n" {
		t.Fatal("Invalid error message")
	}
}
