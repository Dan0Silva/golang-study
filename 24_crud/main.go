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

	fmt.Println("Listening on port 5000 . . .")
	log.Fatal(http.ListenAndServe(":5000", router))
}
