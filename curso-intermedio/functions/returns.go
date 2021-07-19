package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, number := range values {
		total += number
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func main() {
	fmt.Println("Suma", sum(1, 2, 3))
	fmt.Println("Suma", sum(1, 2, 3, 5, 9))
	printNames("Marcos", "Carlos", "Maria", "Carla")
	fmt.Println(getValues(2))
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}
