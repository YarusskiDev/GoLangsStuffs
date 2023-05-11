package banco

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb" // drive de conexão
)

type usuario struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func Conectar() (*sql.DB, error) {
	stringConexao := "Server=DESKTOP-LG2IF4G;Database=Golang;Trusted_Connection=True;"

	db, erro := sql.Open("sqlserver", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	/* fmt.Printf("de? %s", db) */
	return db, nil
}
