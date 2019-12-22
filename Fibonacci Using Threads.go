package main

import (
	"fmt"
)

func main() {
	limit := 12 // number of fibonacci sequences to generate
	jobs := make(chan int, limit) // jobs is a queue of numbers
	result := make(chan int, limit)

	go worker(jobs, result) // activate go workers and see the difference
	// go worker(jobs, result)
	// go worker(jobs, result)
	// go worker(jobs, result)
	// go worker(jobs, result)
	// go worker(jobs, result)

	for i := 0; i < limit; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < limit; j++ {
		fmt.Println(<-result)
	}

}

func worker(jobs <-chan int, result chan<- int) { 
	for n := range jobs { //job : channel used to receive values, result: used to send values
		result <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
