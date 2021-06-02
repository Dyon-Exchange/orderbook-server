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

func TestCancelOrderHandler(t *testing.T) {
	AddAsset("string123")
	Assets["string123"].ProcessLimitOrder(orderbook.Buy, "1", decimal.New(10, 0), decimal.New(10, 0))

	body := CancelOrderRequest{
		Asset:   "string123",
		OrderId: "1",
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/cancelOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CancelOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCancelOrderHandlerWrongMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/cancelOrder", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CancelOrderHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestCancelOrderHandlerEmptyBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/cancelOrder", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CancelOrderHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestCancelOrderHandlerNoAsset(t *testing.T) {
	body := CancelOrderRequest{
		Asset:   "string4321",
		OrderId: "1",
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/cancelOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CancelOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestCancelOrderHandlerGarbageJson(t *testing.T) {
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode("abc")

	req, err := http.NewRequest("POST", "/cancelOrder", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CancelOrderHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Got wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "Malformed JSON.\n" {
		t.Fatal("Invalid error message")
	}
}
