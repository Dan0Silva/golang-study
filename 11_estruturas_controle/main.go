package main

import "fmt"

func sumUpTo(number int) int {
	// for loop

	var sum int

	for aux_number := 1; aux_number <= number; aux_number++ {
		sum += aux_number
	}

	return sum
}

func iterate() {
	// for loop iterate

	names := [3]string{"Joao", "Danilo", "Fabio"}
	for indice, value := range names {
		fmt.Printf("value %s on index %d\n", value, indice)
	}
}

func dayWeek(number int) string {

	// switch

	switch number {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid number"
	}
}

func main() {
	number := 10

	// if else

	if number > 10 {
		fmt.Println("é maior que 10")
	} else {
		fmt.Println("é menor ou igual a 10")
	}

	if other_number := number; other_number > 0 {
		fmt.Println("maior que")
	} else {
		fmt.Println("menor que")
	}

	fmt.Println("day week: ", dayWeek(1))
	fmt.Println("==============================")

	fmt.Printf("the sum of numbers from 1 to %d is: %d \n", number, sumUpTo(number))

	iterate()
}
