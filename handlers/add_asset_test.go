package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddAssetHandler(t *testing.T) {
	body := AddAssetRequest{
		Asset: "string123",
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(body)

	req, err := http.NewRequest("POST", "/addAsset", payload)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddAssetHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAddAssetHandlerEmptyBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/addAsset", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddAssetHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "Body is empty.\n" {
		t.Fatal("Incorrect error message")
	}
}

func TestAddAssetHandlerMalformedJSON(t *testing.T) {
	req, err := http.NewRequest("POST", "/addAsset", bytes.NewBufferString("abcdefgh"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddAssetHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
	if rr.Body.String() != "Malformed JSON.\n" {
		t.Fatal("Incorrect error message")
	}
}

func TestAddAssetHandlerWrongMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/addAsset", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddAssetHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
	if rr.Body.String() != "Method not allowed.\n" {
		t.Fatal("Incorrect error message")
	}
}
