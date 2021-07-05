package main

import "fmt"

func main() {
	// Println
	helloMessage := "Hello"
	worldMessage := "World"
	fmt.Println(helloMessage, worldMessage)

	// Printfn
	nombre := "Platzi"
	cursos := 500
	//%s string, %d number, %v sin saber tipo
	fmt.Printf("%s tiene mas de %d cursos \n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de datos
	fmt.Printf("Hello %T \n", helloMessage)
}
