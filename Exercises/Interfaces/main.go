package main

import "fmt"

// Interface Vechicle
type Vechicle interface {
	Start() string
}

// Car Struct
type Car struct {
	Brand string
}

func (c Car) Start() string {
	return fmt.Sprintf("The car %s is starting!", c.Brand)
}

// Motorcycle Struct
type Motorcycle struct {
	Brand string
}

func (m Motorcycle) Start() string {
	return fmt.Sprintf("The motorcycle %s is starting!", m.Brand)
}

func StartVehicle(v Vechicle) {
	fmt.Println(v.Start())
}

func main() {
	car := Car{Brand: "Kia"}
	motorcycle := Motorcycle{Brand: "Ducati"}
	StartVehicle(car)
	StartVehicle(motorcycle)
}
