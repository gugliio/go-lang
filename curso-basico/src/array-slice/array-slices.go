package main

import "fmt"

func main() {
	var array [4]int
	fmt.Println(array)
	array[0] = 1
	array[1] = 2

	fmt.Println(array, len(array), cap(array))

	// Slices
	slices := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slices, len(slices), cap(slices))

	//Metodos de Slices
	fmt.Println(slices[0])
	fmt.Println(slices[:3])
	fmt.Println(slices[2:4])
	fmt.Println(slices[4:])

	//Append
	slices = append(slices, 8)
	fmt.Println(slices)

	//Append nueva lista
	newSlice := []int{8, 9, 10}
	slices = append(slices, newSlice...)
	fmt.Println(slices)
}
