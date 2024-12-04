package main

import (
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
	dogInJson := `{"name":"petrus","race":"poddle","age":4}`
	var dogInStruct dog // aceita map também

	if error := json.Unmarshal([]byte(dogInJson), &dogInStruct); error != nil {
		log.Fatal(error)
	}

	fmt.Println(dogInStruct)
}
