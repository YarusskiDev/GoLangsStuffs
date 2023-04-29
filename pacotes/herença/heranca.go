package main

import "fmt"

type pessoa struct {
	nome  string
	idade int8
	endereco
}

type endereco struct {
	rua    string
	numero int8
}

func main() {

	usuarioEndereco := endereco{"rafael gonçalves", 102}
	usuarioPessoa := pessoa{"cristiano", 31, usuarioEndereco}

	fmt.Println(usuarioPessoa)

}