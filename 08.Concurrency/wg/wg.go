package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}
	wg.Wait()
}

/***************************************************************
*                      FUNCTIONS                               *
****************************************************************/
func doSomething(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting goroutine %d\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Gorountine %d done!\n", id)
}
