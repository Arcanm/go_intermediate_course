package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {

	ftEmployee := FullTimeEmployee{}
	ftEmployee.Name = "Andres"
	ftEmployee.Age = 29
	ftEmployee.Id = 1
	fmt.Println(ftEmployee)
	fmt.Printf("%v", ftEmployee)
	fmt.Println()
	fmt.Printf("%+v", ftEmployee)
}
