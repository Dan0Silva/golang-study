package main

import "fmt"

func main() {
	user := map[string]string{
		"nome":      "Joao",
		"sobrenome": "Santos",
	}

	courses := map[string]map[string][]string{
		"udemy": {
			"name": []string{"Golang", "Javascript", "Cypress"},
		},
	}

	fmt.Println(user)
	fmt.Println(courses)
}
