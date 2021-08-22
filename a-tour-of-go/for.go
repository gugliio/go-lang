package main

import "fmt"

func main() {
	//"For statement"
	sum := 1
	for i := 0; sum < 1000; i++ {
		sum += sum
	}
	fmt.Println("Total", sum)

	// "While"
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println("Total", sum2)
}
