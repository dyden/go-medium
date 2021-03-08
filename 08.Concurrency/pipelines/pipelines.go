package main

import "fmt"

func main() {
	generator := make(chan int)
	double := make(chan int)
	go generatorNumbers(generator)
	go doubleNumber(generator, double)
	print(double)
}

/***************************************************************
*                      FUNCTIONS                               *
****************************************************************/
func generatorNumbers(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func doubleNumber(input <-chan int, output chan<- int) {
	for i := range input {
		output <- i * 2
	}
	close(output)
}

func print(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
