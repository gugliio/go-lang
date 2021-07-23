package main

import "fmt"

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker id: %d started fib for %d \n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id: %d and job: %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	workers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < workers; i++ {
		go Worker(i, jobs, results)
	}

	for _, number := range tasks {
		jobs <- number
	}

	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}
