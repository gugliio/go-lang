package main

import "fmt"

func main() {
	/*
			defer fmt.Println("mundo")
		   	fmt.Println("Hola")
	*/

	fmt.Println("continua")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Termino")
}
