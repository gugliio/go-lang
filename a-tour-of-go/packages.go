package main

import (
	"fmt"
	"math"
	"math/rand"
)

// func add(a int, b int) == func add(a, b int)
func add(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return a, b
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func main() {
	fmt.Println("Random number using math/rand package", rand.Intn(100))
	fmt.Printf("Raiz cuadrada de un numero %g test\n", math.Sqrt(49))

	//Exported name begins with capital letter (math.Pi)
	fmt.Println("Pi number", math.Pi)

	fmt.Println("Funcion con parametros no exportada", add(42, 13))

	// Funcion con multiples resultados
	a, b := swap("Hola", "Ariel")
	fmt.Println("Funcion con multiples resultados", a, b)

	//Funcion con nombre de variables ya definidas
	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)
}
