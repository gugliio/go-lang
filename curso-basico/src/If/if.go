package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es uno")
	} else {
		fmt.Println("No es uno")
	}

	// With and

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("No es verdad")
	}

	// With or

	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad OR")
	}
}
