package main

import (
	"fmt"
)

func main() {

	var nome string = "Cristiano"
	var numero int = 15
	var ponteiro *int
	ponteiro = &numero

	fmt.Println(nome)
	fmt.Println(numero)
	fmt.Println(ponteiro)
}
