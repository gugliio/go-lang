package main

import "fmt"

func formatter() {
	name := "Andrei"
	fmt.Println("My name is:", name)

	a, b := 4, 6
	fmt.Println("Sum", a+b, "Mean Value:", (a+b)/2)

	fmt.Printf("Your age is: %d\n", 21)

	fmt.Printf("X is: %d y is %f", 5, 6.8)

	fmt.Printf("He says: \"Hello GO!\"")

	figure := "Circle"
	radius := 5
	pi := 3.14

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("PI constant is %f\n", pi)

	fmt.Printf("The diameter of a %s with a Radius of %d is %f\n", figure, radius, float64(pi))

	// Quoted string
	fmt.Printf("Quoted string %q\n", figure)

	// Any Type
	fmt.Printf("The diameter of a %v with a Radius of %v is %v\n", figure, radius, float64(pi))

	// Print type
	fmt.Printf("Radius is type %T\n", radius)
}
