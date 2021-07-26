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

	/* 	for i := 0; i < len(listaTareas.tasks); i++ {
		fmt.Println("Index", i, "Nombre", listaTareas.tasks[i].Nombre)
	} */

	/* for index, tarea := range listaTareas.tasks {
		fmt.Println("Index", index, "Nombre", tarea.Nombre)
	} */

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
