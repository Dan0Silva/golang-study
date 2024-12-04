package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  uint8  `json:"age"`
}

func main() {
	myDog := dog{"petrus", "poddle", 4}

	myDogInJson, error := json.Marshal(myDog)
	if error != nil {
		log.Fatal("convert to json apresenst a strange error")
	}

	fmt.Printf("struct mode: %v\n", myDog)
	fmt.Printf("===============================\n\n")

	fmt.Printf("bytes mode: %v\n", myDogInJson)
	fmt.Printf("===============================\n\n")

	fmt.Printf("json mode: %s\n", bytes.NewBuffer(myDogInJson))
	fmt.Printf("===============================\n\n")
}
