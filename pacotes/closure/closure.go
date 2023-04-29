package main

import "fmt"

func closure() func() {
	texto := "testando closure"

	funcaoClosure := func() {
		fmt.Println(texto)
	}
	return funcaoClosure
}

func main() {

	teste := closure()
	teste()
}
