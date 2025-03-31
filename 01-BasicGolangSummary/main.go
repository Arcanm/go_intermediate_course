package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	// INSTANCE VARIABLES //
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	// CATCH ERRORS //
	myValue, err := strconv.ParseInt("7", 0, 54)
	if err != nil {
		// Ways to print errors
		// fmt.Println(err)
		// fmt.Printf("%v", err)
		log.Fatal(err)
	} else {
		fmt.Println(myValue)
	}

	// SLICES AND MAPS //
	// Map
	m := make(map[string]int)
	m["keyOne"] = 6
	m["keyTwo"] = 10
	fmt.Println(m["keyOne"])

	// Slice
	s := []int{1, 2, 3}

	// Add value to slice
	s = append(s, 16)

	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	// USE GO ROUTINES AND CHANNELS//
	// Create a channel for monitoring Goroutines
	channel := make(chan int)
	// Execute in another routine of main
	go doSomething(channel)
	// Wait for the exit of channel to continue main routine
	<-channel

	// POINTERS //
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println("Memory address:", h, "Value:", *h)
}

// FUNCTIONS //
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
