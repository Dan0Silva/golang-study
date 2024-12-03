package shape

import "math"

type Rectangle struct {
	Height int
	Width  int
}

type Circle struct {
	Radius int
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	result := float64(r.Height) * float64(r.Width)
	result = roundToTwoDecimalPlaces(result)

	return result
}

func (c Circle) Area() float64 {
	result := 3.14159 * (float64(c.Radius) * float64(c.Radius))
	result = roundToTwoDecimalPlaces(result)

	return result
}

func roundToTwoDecimalPlaces(number float64) float64 {
	return math.Round(number*100) / 100
}

// func WriteArea(s Shape) {
// 	fmt.Printf("the area is: %0.2f\n", s.Area())
// }
