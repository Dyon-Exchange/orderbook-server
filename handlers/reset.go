package handlers

import (
	"net/http"
)

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	Reset()
	w.WriteHeader(http.StatusOK)
}
