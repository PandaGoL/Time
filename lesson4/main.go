package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PandaGoL/Time.git/models/"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var employee models.Employee
		err := json.NewDecoder(r.Body).Decode(&employee)
	}
}

func main() {
	handler := &Handler()
	http.Handle("/", handler)

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	srv.ListenAndServe()
}
