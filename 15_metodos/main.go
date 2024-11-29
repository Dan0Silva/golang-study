package main

import "fmt"

type user struct {
	name string
	age  int8
}

func (u user) walking() {
	fmt.Println("i'm walking now...")
}

func (u *user) birthday() {
	u.age++
}

func main() {
	user := user{
		name: "Aoto",
		age:  21,
	}

	user.walking()
}
