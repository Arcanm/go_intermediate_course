package main

import (
	"fmt"
	"sync"
	"time"
)

// doSomething simulates a task that takes 5 seconds to complete
// Parameters:
//   - i: task identifier number
//   - wg: pointer to WaitGroup to signal task completion
//   - ch: buffered channel acting as semaphore to limit concurrent execution
func doSomething(i int, wg *sync.WaitGroup, ch chan int) {
	// Decrement WaitGroup counter when function returns
	defer wg.Done()

	// Print start message with task ID
	fmt.Println("Starting", i)

	// Simulate work by sleeping for 5 seconds
	time.Sleep(5 * time.Second)

	// Print completion message with task ID
	fmt.Println("Ending", i)

	// Release semaphore slot by receiving from channel
	<-ch
}

func main() {
	// Create a buffered channel with capacity 2 that will act as our semaphore
	// This means only 2 goroutines can be running concurrently
	ch := make(chan int, 2)

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launch 10 goroutines
	for i := range 10 {
		// Send value to channel before launching goroutine
		// If channel is full (2 goroutines running), this will block
		// until a slot becomes available
		ch <- 1

		// Add 1 to WaitGroup counter before launching goroutine
		wg.Add(1)

		// Launch goroutine with current index, WaitGroup pointer and channel
		// The goroutine will receive the channel and release a slot when done
		go doSomething(i, &wg, ch)
	}

	// Wait for all goroutines to finish by blocking until WaitGroup counter is 0
	wg.Wait()

	// Close the channel to signal that no more values will be sent
	close(ch)

	// Print completion message
	fmt.Println("Finished")
}
