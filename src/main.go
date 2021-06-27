package main

import "fmt"

func main() {
	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma", result)

	// Resta
	result = y - x
	fmt.Println("Resta", result)

	// Multiplicar
	result = x * y
	fmt.Println("Multiplicar", result)

	// Division
	result = y / x
	fmt.Println("Division", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo", result)

	// Incremental
	x++
	fmt.Println("Incremental", x)

	// Decremento
	x--
	fmt.Println("Decremento", x)

	//Area rectangulo
	baseRectangulo := 10
	alturaRectangulo := 15
	fmt.Println("Area Rect", baseRectangulo*alturaRectangulo)

	//Area trapecio
	baseUno := 10
	baseDos := 30
	altura := 50

	fmt.Println("Area Trapecio", ((baseUno+baseDos)*altura)/2)

	//Area circulo
	const PI = 3.14
	var radio float64 = 10
	fmt.Println("Area Circulo", PI*(radio*radio))
}
