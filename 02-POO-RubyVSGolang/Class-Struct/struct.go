package main

import "fmt"

// Use structs instead classes
type Person struct {
	Name string
}

// Use struct funcs instead class methods
func (p Person) Greetings() string {
	return "Hola, soy " + p.Name
}

func main() {
	p := Person{Name: "Andrés"}
	fmt.Println(p.Greetings())
}
