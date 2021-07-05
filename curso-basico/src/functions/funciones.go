package main

import "fmt"

func main() {
	normalFunction("Hola Mundo")
	tripleArgument(1, 2, "Hola")
	value := returnValue(2)
	fmt.Println("Valor return Value", value)
	value1, value2 := doubleReturn(3)
	fmt.Println("Value 1", value1, "Value 2", value2)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func normalFunction(message string) {
	fmt.Println(message)
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}
