package main

import "fmt"

func recuperaFluxo() {
	if r := recover(); r != nil {
		fmt.Println("fluxo recuperado...")
	}
}

func alunoAprovado(n1, n2 float64) bool {
	/* ele só realmente quebra, se não tiver a função de recover, se tiver recover, ele não permite */
	/* o panic chama todas as linhas de codigo que contenha DEFER */
	defer recuperaFluxo()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é exatamente 6")
}

func main() {
	fmt.Println(alunoAprovado(6, 6))
}
