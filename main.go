package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	}).Methods("GET")

	port := "localhost:5000"
	fmt.Println("server running on " + port)
	http.ListenAndServe(port, r)
}
