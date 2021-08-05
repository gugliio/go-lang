package structure

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radio float64
}

type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}
