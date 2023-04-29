package main

import "fmt"
/* esse DEFER faz com que a linha do codigo tenha um Delay, só vai ser executado depois que tudo for executado
com VOID depois que tudo for executado, com return uma linha antes é executado */

func alunoAprovadoOuNao(n1 int, n2 int) bool {
	defer fmt.Print("Média calculada...")
	fmt.Println("Entrando na função para saber se o aluno foi aprovado...")

	media := (n1 + n2) / 2

	if media > 4 {
		return true
	}
	return false
}
func main() {

	fmt.Println(alunoAprovadoOuNao(1, 9))
}
