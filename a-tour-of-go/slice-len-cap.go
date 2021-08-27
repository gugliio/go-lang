package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlices(s)

	s = s[:0]
	printSlices(s)

	s = s[:4]
	printSlices(s)

	s = s[2:]
	printSlices(s)

	var nilSlice []int
	printSlices(nilSlice)
	if nilSlice == nil {
		fmt.Println("Nil Slice")
	}
}

func printSlices(s []int) {
	fmt.Printf("Len=%d cap=%d %v\n", len(s), cap(s), s)
}
