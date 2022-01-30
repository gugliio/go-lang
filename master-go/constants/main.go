package main

import (
	"fmt"
)

func main() {
	const days int = 74
	fmt.Println(days)

	var i int
	fmt.Println(i)

	const pi float64 = 3.1415
	fmt.Println(pi)
	const secondsInHour = 3600

	duration := 234
	fmt.Printf("Duration in seconds %v\n", duration*secondsInHour)

	x, y := 5, 1
	fmt.Println(x / y)

	const n, m int = 4, 5
	fmt.Println(n, m)
	const n1, m1 = 6, 7
	fmt.Println(n1, m1)

	const (
		min  = -5000
		min2 = -300
		min3 = 100
	)

	fmt.Println(min, min2, min3)

	//Constants rules
	// const temp int = 100
	// temp = 50

	// Not initiate in time on runtime
	// const power = math.Pow(2, 3)

	//Not use a var to initialize a const
	//	t := 5
	//	const tc = t
}
