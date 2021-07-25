package main

import "fmt"

func main() {
	x := 25

	modifyValue(y)
	modifyPointer(y)
}

func holaMundo() {
	fmt.Println("Hola Mundo")
}

func modifyValue(x *int) {
	fmt.Printf("  modifyValue: x=%3d @ %p\n", *x, x)
	*x = 100
	fmt.Printf("  modifyValue: x=%3d @ %p\n", *x, x)
}

func modifyPointer(x *int) {
	fmt.Printf("modifyPointer: x=%3d @ %p\n", *x, x)
	n := 1
	x = &n
	fmt.Printf("modifyPointer: x=%3d @ %p\n", *x, x)
}
