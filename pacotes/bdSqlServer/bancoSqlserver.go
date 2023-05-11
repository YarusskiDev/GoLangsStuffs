package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type usuario struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func main() {

	uriConexao := "Server=DESKTOP-LG2IF4G;Database=goLang;Trusted_Connection=True;"

	db, err := sql.Open("sqlserver", uriConexao)
	if err != nil {
		log.Fatal(err)
	}

	/* linhas, err := db.Query("select * from usuarios") */

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Conexao aberta!")

	statement, err := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if err != nil {
		fmt.Printf("erro de %s", statement)
		return
	}
	//defer statement.Close()

	var usuario usuario
	usuario.Nome = "Cris testando"
	usuario.Email = "teste@email.com"

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		fmt.Print("erro no statement de novo")
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		log.Printf("esse é o erro do statement %s", erro)
		return
	}
	fmt.Printf("o id é %d", idInserido)
}
