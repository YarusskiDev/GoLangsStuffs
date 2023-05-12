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
	router.HandleFunc("/create", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/getall", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/getone/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/update/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/delete/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)
	fmt.Println("ouvindo a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
