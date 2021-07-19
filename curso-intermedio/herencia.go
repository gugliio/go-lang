package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

type PrintInfo interface {
	GetMessage() string
}

func (f FullTimeEmployee) GetMessage() string {
	return "Full Time Employee"
}

func (t TemporaryEmployee) GetMessage() string {
	return "Temporary Time Employee"
}

func getMessage(p PrintInfo) {
	fmt.Println(p.GetMessage())
}

func main() {
	fte := FullTimeEmployee{}
	fte.age = 32
	fte.name = "Ariel"
	fte.id = 1
	fmt.Println(fte)
	tte := TemporaryEmployee{}
	getMessage(fte)
	getMessage(tte)
}
