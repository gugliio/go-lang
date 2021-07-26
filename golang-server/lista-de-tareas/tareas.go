package main

import "fmt"

type tasks struct {
	Nombre      string
	Descripcion string
	Completado  bool
}

type taskList struct {
	tasks []*tasks
}

func (t *taskList) agregarALista(task *tasks) {
	t.tasks = append(t.tasks, task)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
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
	t1 := &tasks{Nombre: "Completar curso Golang", Descripcion: "Completar curso en una semana"}
	t2 := &tasks{Nombre: "Completar curso Python", Descripcion: "Completar curso en una semana"}

	listaTareas := &taskList{
		tasks: []*tasks{t1, t2},
	}

	fmt.Println(listaTareas.tasks[0])

	t3 := &tasks{Nombre: "Completar curso Node Js", Descripcion: "Completar curso en una semana"}
	listaTareas.agregarALista(t3)
	fmt.Println(len(listaTareas.tasks))
	listaTareas.eliminarDeLista(1)
	fmt.Println(len(listaTareas.tasks))
}
