package main

import (
	"advanced_test/shape"
	"fmt"
)

func main() {
	rectangle := shape.Rectangle{Height: 10, Width: 15}
	circle := shape.Circle{Radius: 10}

	fmt.Printf("[ %0.2f | %0.2f ]\n", rectangle.Area(), circle.Area())
}
