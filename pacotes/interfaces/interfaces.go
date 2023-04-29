package main

import "fmt"

type forma interface {
	area() int
}

type retangulo struct {
	largura int
}

type quadrado struct {
	altura int
}

func (r retangulo) area() int {
	return 5 * r.largura
}

func (q quadrado) area() int {
	return 5 * q.altura
}

func Calcula(f forma) {
	fmt.Printf("O valor desse caralho é %d", f.area())
}

func main() {

	quadrado := quadrado{altura: 2}
	retangulo := retangulo{largura: 5}

	Calcula(quadrado)
	Calcula(retangulo)
}
