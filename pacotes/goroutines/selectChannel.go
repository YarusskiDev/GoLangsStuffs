package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 2)
			canal1 <- "canal1 testando"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "canal2 testando saporra"
		}
	}()

	for {
		select {
		case mensagem := <-canal1:
			fmt.Println(mensagem)
		case mensagem2 := <-canal2:
			fmt.Println(mensagem2)
		}
	}

}
