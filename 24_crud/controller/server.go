package controller

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type user struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []user
	var user user

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Erro ao se conectar com o banco de dados."))
		return
	}
	defer db.Close()

	lines, err := db.Query("select * from USERS")
	if err != nil {
		w.Write([]byte("Erro ao buscar usuários."))
	}
	defer lines.Close()

	for lines.Next() {
		if err := lines.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear usuário."))
			return
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Erro ao converter usuários para formato json."))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id := params["id"]
	var user user

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Erro ao se conectar com o banco de dados."))
		return
	}
	defer db.Close()

	line, err := db.Query("select * from USERS where ID = ?", Id)
	if err != nil {
		w.Write([]byte("Erro ao buscar usuário no banco de dados."))
		return
	}

	if line.Next() {
		if err = line.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear usuário."))
			return
		}
	}

	fmt.Println("testing: ", Id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter usuários para formato json."))
	}
}
