package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	handler := Handler{}

	http.Handle("/api", &handler)
	http.ListenAndServe(":8080", nil)
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"message": "hello world",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
