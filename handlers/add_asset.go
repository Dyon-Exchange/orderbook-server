package handlers

import (
	"encoding/json"
	"net/http"
)

type AddAssetRequest struct {
	Asset string
}

func AddAssetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Body is empty.", http.StatusBadRequest)
		return
	}

	var req AddAssetRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Malformed JSON.", http.StatusBadRequest)
		return
	}

	AddAsset(req.Asset)
	w.WriteHeader(http.StatusOK)
}
