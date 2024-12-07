package controller

import (
	"crud/database"
	"encoding/json"
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
		return
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

	_, err = statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
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

	if _, err := uuid.Parse(Id); err != nil {
		w.Write([]byte("ID invalido"))
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter usuários para formato json."))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user user

	Id := params["id"]

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler corpo da requisição."))
		return
	}

	if err = json.Unmarshal(reqBody, &user); err != nil {
		w.Write([]byte("Erro ao converter requisição para json."))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Erro ao tentar conexão com banco"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update USERS set NAME = ?, EMAIL = ? where ID = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Email, Id); err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	Id := params["id"]

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from USERS where ID = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar statement."))
		return
	}
	defer statement.Close()

	if _, err = statement.Exec(Id); err != nil {
		w.Write([]byte("Erro ao deletar usuário."))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
