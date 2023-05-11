package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello GoLang here mada faka"))
	})

	fmt.Println("Ouvindo na porta 5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
