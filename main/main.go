package main

import (
	"net/http"

	"./api/v1/tickets"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	V1 := r.PathPrefix("/v1").Subrouter()

	V1.HandleFunc("/tickets", tickets.GetList).Methods("GET", "POST")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}

}
