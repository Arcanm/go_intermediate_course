package main

import "fmt"

// In Ruby, you use include, while in Go, you just implement
// the required methods, and the interface is recognized automatically.
type Animal interface {
	MakeSound() string
}

type Dog struct{}

func (dog Dog) MakeSound() string {
	return "Guau"
}

func main() {
	var a Animal = Dog{}
	fmt.Println(a.MakeSound())
}
