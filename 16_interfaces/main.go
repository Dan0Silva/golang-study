package main

import "fmt"

// interface

type rectangle struct {
	height int
	width  int
}

type circle struct {
	radius int
}

type shape interface {
	area() float64
}

func (r rectangle) area() float64 {
	return float64(r.height) * float64(r.width)
}

func (c circle) area() float64 {
	return 3.14 * (float64(c.radius) * float64(c.radius))
}

func writeArea(s shape) {
	fmt.Printf("the area is: %0.2f\n", s.area())
}

//interface generica

func generic(s interface{}) {
	fmt.Println(s)
}

func main() {
	rectangle := rectangle{height: 10, width: 15}
	writeArea(rectangle)

	circle := circle{radius: 6}
	writeArea(circle)

	fmt.Println("=======================")

	generic("string")
	generic(12)
	generic(true)
}
