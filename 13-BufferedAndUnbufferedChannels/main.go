package main

import "fmt"

func main() {
	// Unbuffered channel
	// When you send a message to a unbuffered channel,
	// the sender will wait until the receiver is ready
	// to receive the message.
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)

	// Buffered channel
	// When you send a message to a buffered channel,
	// the sender will not wait until the receiver is ready
	// to receive the message.
	c2 := make(chan int, 2)
	c2 <- 2
	c2 <- 3
	fmt.Println(<-c2)
	fmt.Println(<-c2)
}
