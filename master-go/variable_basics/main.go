package main

import "fmt"

func main() {
	// Variables declaration

	var age int = 30
	fmt.Println("Age:", age)

	var name = "Dan"
	_ = name

	s := "Learing golang"
	fmt.Println(s)

	name = "Andrei"
	name1 := "John"
	_ = name1

	// Multiples declaration

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

	// Types and zero values
	sum := 5 + 2.3
	fmt.Println(sum)

	var ab = 4
	var bc = 5.2

	ab = int(bc)
	fmt.Println(ab)

	var value int
	var price float64
	var names string
	var done bool
	fmt.Println(value, price, names, done)

	// comments
	// This is a new comment

	// declaring a new variable of type int
	/*
		 age := 18 // short- declaration - inline deaclaration
		_ = age
	*/

}
