package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	var car2 car
	car2.brand = "Ferrari"
	fmt.Println(car2)
}
