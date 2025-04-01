package main

import "fmt"

// VARIADIC FUNCTIONS, YOU CAN USE ... FOR SEND DYNAMIC PARAMS OF DETERMINATED TYPE
// EXAMPLE: sum func can recieve any numbers or type integer
func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}

	return total
}

func printNames(values ...string) {
	for _, name := range values {
		fmt.Println(name)
	}
}

// YOU CAN DEFINE NAME FOR RETURN VALUES, AND YOU CAN ASSING VALUE INSIDE FUNC
// WHEN YOU USE return AUTOMATICALLY RETURN THE THREE VARIABLES WITH HIS VALUE
func getValues(x int) (double, triple, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))

	printNames("Andres")
	printNames("Andres", "Sergio", "Carlangas")

	fmt.Println(getValues(2))
}
