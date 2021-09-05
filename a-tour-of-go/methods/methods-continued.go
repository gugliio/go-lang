package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (m MyFloat) Abs() float64 {
	if m < 0 {
		return float64(-m)
	}
	return float64(m)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
