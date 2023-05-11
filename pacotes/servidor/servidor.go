package servidor

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	banco "meumodulo/BancoDeDados"

	"net/http"
)

type usuario struct {
	Id    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo"))
		return
	}

	var usuario usuario

	if err := json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("erro ao converter o usuario"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("erro ao conectar ao banco"))
		return
	}
	defer db.Close()

	query := "INSERT INTO dbo.usuarios (nome, email) OUTPUT INSERTED.ID VALUES (@Nome, @Email)"
	// fmt.Println("Query SQL:", query)

	_, erro = db.Exec(query, sql.Named("Nome", usuario.Nome), sql.Named("Email", usuario.Email))
	db.Prepare()
	if erro != nil {
		log.Println("Erro ao inserir usuário:", erro)
		w.Write([]byte("erro ao executar o statement"))
		return
	}

	var lastId int64
	row := db.QueryRowContext(context.Background(), usuario.Nome, usuario.Email)
	erro = row.Scan(&lastId)

	if erro := db.QueryRow("SELECT SCOPE_IDENTITY()").Scan(&lastId); erro != nil {
		log.Println("Erro ao pegar ultimo Id:", erro)
		w.Write([]byte("erro ao pegar ultimo Id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("usuario inserido com sucesso! Id: %d", lastId)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("erro ao conectar ao banco"))
		return
	}
	defer db.Close()
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("erro ao buscar usuarios"))
		return
	}
	defer linhas.Close()

	var usuariosBdArray []usuario
	for linhas.Next() {
		var usuarioBd usuario
		if erro := linhas.Scan(&usuarioBd.Id, &usuarioBd.Nome, &usuarioBd.Email); erro != nil {
			w.Write([]byte("erro ao escanear usuarios"))
			return
		}
		usuariosBdArray = append(usuariosBdArray, usuarioBd)
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuariosBdArray); erro != nil {
		w.Write([]byte("erro ao converter os usuarios para json!"))
	}
}
