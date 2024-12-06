package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type sqlConnection struct {
	name   string
	pass   string
	tcp    string
	dbName string
}

func main() {
	s := sqlConnection{
		name:   "danilo",
		pass:   "2204",
		tcp:    "172.17.0.2:3306",
		dbName: "golang_study",
	}

	stringConnection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", s.name, s.pass, s.tcp, s.dbName)

	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successful connection!")
	}

	lines, err := db.Query("select * from USERS")
	if err != nil {
		log.Fatal(err)
	}
	defer lines.Close()
}
