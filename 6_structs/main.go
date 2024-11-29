package main

import "fmt"

type user struct {
	name string
	age  int8
	mail string
}

func newUser(name, mail string, age int8) *user {
	return &user{
		name: name,
		mail: mail,
		age:  age,
	}
}

func main() {
	profile := newUser("Danilo A. Silva", "test@mail.com", 22)
	fmt.Println(profile.name)
}
