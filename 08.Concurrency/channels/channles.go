package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//UNBUFFERED CHANNELS
	// c1 := make(chan int)
	//BUFFERED CHANNELS
	c2 := make(chan int, 3)
	// c1 <- 1
	//fmt.Println(<-c1)
	c2 <- 1
	c2 <- 2
	c2 <- 3
	fmt.Println(<-c2, <-c2, <-c2)

	//CHANNEL SYNCHRONIZATION
	c3 := make(chan int, 1) //MAX 1 PROCCESS CAN ACCESS THE CHANNEL
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c3 <- 1
		wg.Add(1)
		go doSomething(i, &wg, c3)
	}
	wg.Wait()
}

/***************************************************************
*                      FUNCTIONS                               *
****************************************************************/
func doSomething(id int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Starting goroutine %d\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Gorountine %d done!\n", id)
	<-c

}
