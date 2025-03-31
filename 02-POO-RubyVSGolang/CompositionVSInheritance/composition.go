package main

import "fmt"

type Animal struct{}

func (a Animal) Eat() string {
	return "Eating..."
}

// Go does not support inheritance, instead, it uses struct
// embedding for code reuse.
type Dog struct {
	Animal // Composition, not inheritance
}

func (d Dog) Bark() string {
	return "Woof"
}

func main() {
	d := Dog{}
	fmt.Println(d.Eat()) // Calling method from Animal
	fmt.Println(d.Bark())
}
