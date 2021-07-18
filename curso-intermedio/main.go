package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("---------------------------")

	value, err := strconv.ParseInt("perro", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(value)
	}

	fmt.Println("---------------------------")

	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	fmt.Println("---------------------------")

	slice := []int{1, 2, 3}
	for index, value := range slice {
		fmt.Println("Index", index, "Value", value)
	}

	fmt.Println("---------------------------")

	slice = append(slice, 16)

	for index, value := range slice {
		fmt.Println("Index", index, "Value", value)
	}

	fmt.Println("---------------------------")
}
