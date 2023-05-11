package main

import (
	"fmt"
	"log"
	"meumodulo/servidor"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/criar", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/todos", servidor.BuscarUsuarios).Methods(http.MethodGet)
	fmt.Println("ouvindo a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
