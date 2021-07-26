package main

import "fmt"

type tasks struct {
	Nombre      string
	Descripcion string
	Completado  bool
}

func (t *tasks) marcarCompleta() {
	t.Completado = true
}

func (t *tasks) actualizarDescripcion(desc string) {
	t.Descripcion = desc
}

func (t *tasks) actualizarNombre(nombre string) {
	t.Nombre = nombre
}

func main() {
	t := &tasks{Nombre: "Completar curso", Descripcion: "Completar curso en una semana"}
	fmt.Printf("%+v\n", t)

	t.marcarCompleta()
	t.actualizarNombre("Finalizar curso")
	t.actualizarDescripcion("Finalizar curso lo antes posible")
	fmt.Printf("%+v\n", t)
}
