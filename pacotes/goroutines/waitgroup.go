package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Cristiano")
		waitGroup.Done() //-1
	}()

	go func() {
		escrever("Teste com waitGroup")
		waitGroup.Done() //-1
	}()

	waitGroup.Wait() //0
}
