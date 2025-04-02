package main

import "fmt"

// This example demonstrates the use of directional channels in Go
// Directional channels can be either send-only (chan<-) or receive-only (<-chan)
// This provides type safety and makes the code's intent clearer

// Generator sends values to a channel (ONLY WRITING)
// Parameters:
//   - c: send-only channel (chan<- int) that will receive sequential numbers
//
// The function:
//  1. Generates numbers from 0 to 9
//  2. Sends each number to the channel
//  3. Closes the channel when done
func Generator(c chan<- int) {
	for i := range 10 {
		c <- i
	}
	close(c)
}

// Double receives values from a channel and sends doubled values to another channel
// Parameters:
//   - in: receive-only channel (<-chan int) that provides input values
//   - out: send-only channel (chan<- int) that will receive doubled values
//
// The function:
//  1. Reads values from input channel until it's closed
//  2. Multiplies each value by 2
//  3. Sends the result to output channel
//  4. Closes the output channel when done
func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value * 2
	}
	close(out)
}

// Printer receives and prints values from a channel
// Parameters:
//   - c: receive-only channel (<-chan int) that provides values to print
//
// The function:
//  1. Reads values from the channel until it's closed
//  2. Prints each value to standard output
func Printer(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	// Create two bidirectional channels
	// - generator: will carry the original numbers
	// - doubles: will carry the doubled numbers
	generator := make(chan int)
	doubles := make(chan int)

	// Start Generator goroutine to produce numbers
	go Generator(generator)

	// Start Double goroutine to consume numbers from generator
	// and produce doubled numbers
	go Double(generator, doubles)

	// Print both the original and doubled numbers
	// Note: This will block until both channels are closed
	Printer(generator)
	Printer(doubles)
}
