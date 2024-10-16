package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var routes = map[string]http.Handler{
	"/api": &Handler{},
}

func main() {
	for path, handler := range routes {
		http.Handle(path, handler)
	}
	log.Println("Server started on port 8080")
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
