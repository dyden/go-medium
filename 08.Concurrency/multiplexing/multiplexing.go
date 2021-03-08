package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go doSomething(d1, c1, 1)
	go doSomething(d2, c2, 2)

	// fmt.Println(<-c1, <-c2) //BLOCKING (RESULT: 1 2)

	for i := 0; i < 2; i++ {
		select {
		case c1 := <-c1:
			fmt.Println(c1)
		case c2 := <-c2:
			fmt.Println(c2)
		}
	}
}

/***************************************************************
*                      FUNCTIONS                               *
****************************************************************/
func doSomething(t time.Duration, c chan<- int, param int) {
	time.Sleep(t)
	c <- param
}
