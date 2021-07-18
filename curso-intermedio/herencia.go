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
}

/* func getMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
} */

func main() {
	fte := FullTimeEmployee{}
	fte.age = 32
	fte.name = "Ariel"
	fte.id = 1

	fmt.Println(fte)
	// getMessage(fte)
}
