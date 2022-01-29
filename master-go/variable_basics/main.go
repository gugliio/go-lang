package main

import "fmt"

func main() {
	var age int = 30
	fmt.Println("Age:", age)

	var name = "Dan"
	_ = name

	s := "Learing golang"
	fmt.Println(s)

	name = "Andrei"
	name1 := "John"
	_ = name1

	car, cost := "Audi", 5000
	fmt.Println(car, cost)
	car, year := "BMW", 2006
	_ = year

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    string
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	_, _ = i, j

	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)
}
