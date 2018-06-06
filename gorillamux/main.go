package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mesaje desde el metodo GET")
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mesaje desde el metodo POST")
}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mesaje desde el metodo PUT")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mesaje desde el metodo DELETE")
}

func main() {
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", PostUsers).Methods("POST")
	r.HandleFunc("/api/users", PutUsers).Methods("PUT")
	r.HandleFunc("/api/users", DeleteUsers).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
