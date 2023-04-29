package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	go escrever("Cristiano testando...")
	escrever("testango palavra reservada GO")
}
