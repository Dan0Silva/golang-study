package controller

import (
	"crud/database"
	"encoding/json"
	"io"
	"net/http"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição."))
		return
	}

	var user user
	if err = json.Unmarshal(reqBody, &user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struct."))
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Erro ao tentar conexão com banco"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT
	statement, err := db.Prepare("insert into USERS (NAME, EMAIL) values (?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insertion, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	_, err = insertion.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário cadastrado com sucesso"))

}
