package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)

	e2 := Employee{1, "Ari", true}
	fmt.Printf("%v\n", e2)

	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 1
	e3.name = "Name"
	fmt.Printf("%v\n", *e3)

	e4 := NewEmployee(1, "Ariel", false)
	fmt.Printf("%v\n", *e4)
}
