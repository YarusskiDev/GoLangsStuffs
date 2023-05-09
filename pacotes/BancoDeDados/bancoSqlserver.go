package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {

	uriConexao := "Server=DESKTOP-LG2IF4G;Database=Golang;Trusted_Connection=True;"

	db, err := sql.Open("sqlserver", uriConexao)
	if err != nil {
		log.Fatal(err)
	}

	linhas, err := db.Query("select * from usuarios")

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Conexao aberta!")
}
