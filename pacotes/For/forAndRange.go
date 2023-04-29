package main

import "fmt"

func main() {

	slice := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Println("esse é o valor", i)
	}

	for indice, valor := range slice{
		fmt.Println(indice, valor)
	}
}
