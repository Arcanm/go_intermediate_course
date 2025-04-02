package main

import (
	"fmt"
	"time"
)

// DoSomething simulates a task that takes a specified duration to complete
// Parameters:
//   - i: Duration to sleep/wait
//   - c: Send-only channel to send result
//   - param: Value to send on the channel after sleeping
//
// The function:
//  1. Sleeps for the specified duration
//  2. Sends the param value on the channel
func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}

func main() {
	// Create two unbuffered channels for communication
	c1 := make(chan int)
	c2 := make(chan int)

	// Define different durations for our goroutines
	d1 := 6 * time.Second // Longer duration
	d2 := 2 * time.Second // Shorter duration

	// Launch two goroutines with different sleep durations
	// First goroutine will take 6 seconds
	go DoSomething(d1, c1, 1)
	// Second goroutine will take 2 seconds
	go DoSomething(d2, c2, 2)

	// The commented code below shows the blocking nature of channels
	// When reading directly, it will always read from c1 first, then c2,
	// regardless of which goroutine finishes first
	// fmt.Println(<-c1)
	// fmt.Println(<-c2)

	// Using select to handle multiple channels
	// This will read from whichever channel has data ready first
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1) // Will print 1 (after 6 seconds)
		case msg2 := <-c2:
			fmt.Println(msg2) // Will print 2 (after 2 seconds)
		}
	}
}
