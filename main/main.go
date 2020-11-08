package main

import (
	"fmt"
	"net/http"

	"./api/v1/tickets"

	"github.com/gorilla/mux"
)

func methodNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if _, err := fmt.Fprintln(w, "123"); err != nil {
		return
	}
}

func main() {

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(methodNotFound)

	V1 := r.PathPrefix("/v1").Subrouter()

	V1.HandleFunc("/tickets", tickets.GetList).Methods("GET", "POST")
	V1.HandleFunc("/ticket/create", tickets.Create).Methods("GET", "POST")
	V1.HandleFunc("/ticket/add", tickets.Add).Methods("GET", "POST")
	V1.HandleFunc("/ticket/{id:[0-9]+}", tickets.Read).Methods("GET", "POST")

	if err := http.ListenAndServe(":8080", r); err != nil {
		return
	}

}
