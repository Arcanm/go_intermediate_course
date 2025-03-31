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
	TaxRate int
}

type TemporaryEmployee struct {
	Person
	Employee
	EndDate string
}

// IMPLEMENTING INTERFACE FOR FULLTIME AND TEMPORARY EMPLOYEE WITH HIS IMPLEMENTATION
type PrintInfo interface {
	GetMessage() string
}

func (ftEmployee FullTimeEmployee) GetMessage() string {
	return fmt.Sprintf("\nFull Time Employee\nId: %d Name: %s Age: %d Tax Rate: %d",
		ftEmployee.Id,
		ftEmployee.Name,
		ftEmployee.Age,
		ftEmployee.TaxRate)
}

func (tEmployee TemporaryEmployee) GetMessage() string {
	return fmt.Sprintf("\nTemporary Employee\nId: %d Name: %s age: %d End Date: %s",
		tEmployee.Id,
		tEmployee.Name,
		tEmployee.Age,
		tEmployee.EndDate,
	)
}

// Since GetMessage receives a parameter of type PrintInfo, it can accept any type that has the GetMessage() method.
func GetMessage(p PrintInfo) {
	fmt.Println(p.GetMessage())
}

//

func main() {

	ftEmployee := FullTimeEmployee{}
	ftEmployee.Name = "Andres"
	ftEmployee.Age = 29
	ftEmployee.Id = 1

	temporaryEmployee := TemporaryEmployee{}
	ftEmployee.Name = "PC"
	ftEmployee.Age = 25
	ftEmployee.Id = 2

	GetMessage(ftEmployee)
	GetMessage(temporaryEmployee)
}
