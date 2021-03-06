package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/reset", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ResetHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
