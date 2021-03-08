package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 2
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}
	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)
	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

/***************************************************************
*                      FUNCTIONS                               *
****************************************************************/
func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d is working on job %d\n", id, job)
		fib := Fibonaaci(job)
		fmt.Printf("Worker with id %d finished job %d and the result is %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonaaci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonaaci(n-1) + Fibonaaci(n-2)

}
