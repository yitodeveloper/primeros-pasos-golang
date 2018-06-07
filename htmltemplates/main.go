package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func login(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "Página no encontrada")
	} else {
		temp.Execute(w, nil)
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", login)

	server := &http.Server{
		Addr:           ":4040",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
