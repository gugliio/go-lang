package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale2(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("First ABS", v.Abs())
	v.Scale(10)
	fmt.Println("Second ABS", v.Abs())

	v2 := Vertex{3, 4}
	fmt.Println("First ABS", v2.Abs())
	Scale2(&v2, 10)
	fmt.Println("Second ABS", v2.Abs())

	v3 := Vertex{3, 4}
	fmt.Println(v3.Abs())
	fmt.Println(AbsFunc(v3))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
