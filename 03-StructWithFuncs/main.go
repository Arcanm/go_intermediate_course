package main

import "fmt"

type Employee struct {
	Id   int
	Name string
}

// If you want to modify a value of your struct, you need to pass as pointer
// If you dont use *, the method works with a copy and doesn modify the values
// that you want to modify.
func (e *Employee) SetId(id int) {
	e.Id = id
}

func (e *Employee) SetName(name string) {
	e.Name = name
}

func (e Employee) String() string {
	return fmt.Sprintf("The ID of employee is %d, and his name is %s", e.Id, e.Name)
}

func main() {
	e := Employee{Id: 1, Name: "Andres"}
	// Print formatted struct defined in method String() -> Stringer
	fmt.Println(e)
	// If stringer was not defined next line
	// prints field name and value of the struct
	// fmt.Printf("%+v\n", e)

	e.Id = 2
	e.Name = "Paola"
	e.SetId(6)
	fmt.Println(e)
	e.SetName("JC")
	fmt.Println(e)
}
