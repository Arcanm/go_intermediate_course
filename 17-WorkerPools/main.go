package main

import (
	"fmt"
	"time"
)

// Worker represents a goroutine that processes jobs from a jobs channel and sends results to a results channel
// Parameters:
//   - id: unique identifier for the worker
//   - jobs: receive-only channel that provides Fibonacci numbers to calculate
//   - results: send-only channel where calculated Fibonacci results are sent
//
// The function:
//  1. Announces when the worker starts
//  2. Continuously reads jobs from the jobs channel until it's closed
//  3. For each job, calculates the Fibonacci number and sends it to results
//  4. Prints progress messages for each job
func Worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Println("Worker with id", id, "started")
	for j := range jobs {
		fmt.Printf("Worker with id %v started job %v\n", id, j)
		results <- Fibonacci(j)
		fmt.Printf("Worker with id %v finished job %v\n", id, j)
	}
}

// Fibonacci calculates the nth number in the Fibonacci sequence recursively
// Parameters:
//   - n: position in Fibonacci sequence to calculate
//
// Returns:
//   - The nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	// Define a slice of tasks (positions in Fibonacci sequence to calculate)
	tasks := []int{7, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42}

	// Create buffered channels for jobs and results
	// Buffer size matches number of tasks to prevent blocking
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	// Start 3 worker goroutines
	// Each worker will process jobs from the jobs channel
	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results)
	}

	// Wait for workers to initialize
	time.Sleep(3 * time.Second)

	// Send all tasks to the jobs channel
	// Add 1 second delay between tasks to simulate real-world job distribution
	for _, task := range tasks {
		println("Sending task", task)
		jobs <- task
		time.Sleep(1 * time.Second)
	}

	// Close jobs channel to signal no more tasks
	close(jobs)

	// Collect and print results
	// Number of results to collect matches number of tasks
	for a := 1; a <= len(tasks); a++ {
		fmt.Println(<-results)
	}
}
