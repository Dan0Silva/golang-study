package main

import "fmt"

type people struct {
	name string
	age  int8
}

// "herança", só pra dizer que tem
type student struct {
	people
	course  string
	college string
}

func main() {

	p1 := people{"Arthur Morgan", 33}
	fmt.Println(p1)

	st1 := student{p1, "Software Engineering", "Harvard"}
	fmt.Println(st1)

	st2 := student{people{"John Marston", 27}, "software engineering", "harvard"}
	fmt.Println(st2)
}
