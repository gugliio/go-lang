package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	fmt.Println("-----------------------------")
	e := Employee{}
	fmt.Printf("%v\n", e)
	fmt.Println("-----------------------------")
	e.id = 1
	e.name = "Ariel"
	fmt.Printf("%v\n", e)
	fmt.Println("-----------------------------")
	e.SetId(5)
	e.SetName("Marcos")
	fmt.Printf("%v\n", e)
	fmt.Println("-----------------------------")
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
