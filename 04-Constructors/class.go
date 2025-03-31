package main

import "fmt"

type Employee struct {
	Id       int
	Name     string
	Vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		Id:       id,
		Name:     name,
		Vacation: vacation,
	}
}

func main() {
	// Initialize a struct (class)
	// With zero values, take care with zero values is not nulls/nils
	e := Employee{}
	fmt.Println(e)

	// With initial values
	e2 := Employee{
		Id:       1,
		Name:     "Andres",
		Vacation: true,
	}
	fmt.Println(e2)

	// With reserved word new, its similar to first method
	// It returns the reference (&) not the values
	// You can initialize Employee without params
	e3 := new(Employee)
	fmt.Println(e3)
	fmt.Println(*e3)

	// BEST WAY TO INITIALIZE A STRUCT (class)
	// Custom constructor(initializer) for struct Employee
	// You need initialize struct with all the params
	e4 := NewEmployee(1, "Andrew", true)
	fmt.Println(e4)
	fmt.Println(*e4)

}
