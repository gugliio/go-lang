package main

import "fmt"

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

type figuras2D interface {
	area() float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println(f.area())
}

func main() {
	myCuadrado := cuadrado{2}
	calcular(myCuadrado)

	myRectangulo := rectangulo{2, 3}
	calcular(myRectangulo)

	// Lista de interfaces

	myInferface := []interface{}{"hola", 12, 4.5}
	fmt.Println(myInferface...)
}
