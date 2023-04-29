package main

import "fmt"

func main() {
	canal := make(chan string, 2) // o parametro 2,da capacidade ao canal para funcionar fora de funções
	canal <- "Olá mundo"          // ele não bloqueia quando faz o envio de mensagem, quando tem buffer
	canal <- "bang xarope da porra"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
