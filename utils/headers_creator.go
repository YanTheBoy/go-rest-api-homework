package utils

import "net/http"

func SetUpHeaders(statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	switch statusCode {
	case http.StatusOK:
		w.WriteHeader(http.StatusOK)
	case http.StatusCreated:
		w.WriteHeader(http.StatusCreated)
	}
}
