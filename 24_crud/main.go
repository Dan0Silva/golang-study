package main

import (
	"crud/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", controller.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", controller.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controller.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", controller.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", controller.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Listening on port 5000 . . .")
	log.Fatal(http.ListenAndServe(":5000", router))
}
