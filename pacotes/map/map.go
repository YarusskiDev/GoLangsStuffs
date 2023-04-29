package main

import (
	"fmt"
)

func main() {
	meuMap := map[string]string{"nome": "Cristiano",
		"age": "20",
	}

	meuMap2 := make(map[string]int)
	fmt.Println(meuMap2)

	fmt.Println(meuMap)
	meuMap["sobrenome"] = "rodrigues de souza"
	fmt.Println(meuMap["sobrenome"])

}
