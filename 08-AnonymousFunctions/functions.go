package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	fmt.Println(x)

	// WITH INMEDIATLY EXECUTION -> USE ( ) AT THE END
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)

	// WITHOUT INMEDIATLY EXECUTION -> USE ( ) AT THE END
	double := func() int {
		x *= 2
		return x
	}
	// EXECUTE VARIABLE WITH THE FUNCTION
	fmt.Println(double())
	fmt.Println(double())

	// WITH GO ROUTINES
	c := make(chan int)
	go func() {
		fmt.Println("Initializing go routine")
		time.Sleep(5 * time.Second)
		fmt.Println("Ending of Goroutine")
		c <- 1
	}()
	<-c
}
