package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Random number using math/rand package", rand.Intn(100))
	fmt.Printf("Raiz cuadrada de un numero %g test\n", math.Sqrt(49))

	//Exported name begins with capital letter (math.Pi)
	fmt.Println("Pi number", math.Pi)
}
