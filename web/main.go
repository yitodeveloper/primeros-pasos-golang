package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func registerAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Registro Bakan</h1><a href=\" / \">Volver</a>")
}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

func main() {

	msg := mensaje{
		msg: "Holitas",
	}

	mux := http.NewServeMux() // Enrutador del paquete http

	/****** Servir Archivos Estaticos ******/
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	/****** Fin Servir Archivos Estaticos ******/

	mux.HandleFunc("/vc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Dopleye</h1><a href=\" /registro \">Registrate</a>")
	})
	mux.HandleFunc("/registro", registerAction)
	mux.Handle("/hola", msg)

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Dopleye</h1><a href=\" /registro \">Registrate</a>")
	})

	http.HandleFunc("/registro", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Registro Bakan</h1><a href=\" / \">Volver</a>")
	})*/

	/*http.ListenAndServe(":8080", mux)*/ // Servidor Basico

	//Estructura server customizada
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Escuchando ...")

	server.ListenAndServe()
}
