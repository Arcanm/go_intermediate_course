package main

import (
	"fmt"
	"sync"
	"time"
)

// WAITGROUP
// WaitGroup is used to wait for a group of goroutines to finish.
// It is a struct that has a counter and a channel.
// The counter is the number of goroutines to wait for.
// The channel is used to send the signal to the goroutines to finish.

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Ending", i)
}

func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go doSomething(i, &wg)
	}
	wg.Wait()
	fmt.Println("Finished")
}
