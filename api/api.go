package api

import (
	"encoding/json"
	"net/http"
)

// Order params
type OrderParams struct {
	Username string
}

// Order response
type OrderResponse struct {
	// Success code, 200
	Code int

	// Order details
	OrderId string
	Product string
	Quantity int64
}

// Error response
type Error struct {
	Code int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error {
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Unexpected Error Occured", http.StatusInternalServerError)	
	}
)
