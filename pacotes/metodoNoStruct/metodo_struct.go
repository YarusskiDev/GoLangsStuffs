package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (user usuario) mostrarProps() {
	fmt.Println(user.idade)
}

/* pra aumentar o valor, preciso usar o ponteiro com * */
func (user *usuario) aumentaIdade() {
	user.idade += 2
}

func main() {
	usuario1 := usuario{nome: "Cristiano", idade: 31}
	usuario1.mostrarProps()
	usuario1.aumentaIdade()
	usuario1.mostrarProps()
}
