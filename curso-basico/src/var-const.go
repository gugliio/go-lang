package main

import "fmt"

func main() {
	// Constantes
	const PI float64 = 3.14
	const PI2 = 3.1415

	fmt.Println("Pi", PI)
	fmt.Println("Pi", PI2)

	//Var Int
	base := 12
	var altura int = 14
	var area int

	fmt.Println("Base", base, "Altura", altura, "Area", area)

	//Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area de un cuadrado
	const BASECUADRADO = 10
	areaCuadrado := BASECUADRADO * BASECUADRADO
	fmt.Println("Area Cuadrado", areaCuadrado)
}
