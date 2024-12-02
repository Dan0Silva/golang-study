package main

import "fmt"

/*
	imagine uma fila de tarefas, onde temos varios processos
	pegando tarefas dessa fila para fazer de maneira
	independente.
*/

func fibonacci(position int) int {

	if position <= 2 {
		return 1
	}

	x := fibonacci(position-1) + fibonacci(position-2)
	return x
}

func worker(tasks <-chan int, results chan<- int) {
	for task := range tasks {
		results <- fibonacci(task)
	}
}

func main() {
	tasks := make(chan int, 50)
	results := make(chan int, 50)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for num := 1; num <= 45; num++ {
		tasks <- num
	}
	close(tasks)

	if i := 0; true {
		for result := range results {
			fmt.Printf("p[%d]: %d\n", i, result)
			i++
		}
	}
}
