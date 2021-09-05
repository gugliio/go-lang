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

// The first is so that the method can modify the value that its receiver points to.
// The second is to avoid copying the value on each method call.
//This can be more efficient if the receiver is a large struct, for example.

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

	v4 := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v4, v4.Abs())
	v4.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v4, v4.Abs())
}
