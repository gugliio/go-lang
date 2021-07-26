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

func (t taskList) imprimirLista() {
	for _, task := range t.tasks {
		fmt.Println("Nombre", task.Nombre)
		fmt.Println("Descripcion", task.Descripcion)
	}
}

func (t taskList) imprimirListaCompletados() {
	for _, task := range t.tasks {
		if task.Completado == true {
			fmt.Println("Nombre", task.Nombre)
			fmt.Println("Descripcion", task.Descripcion)
		}
	}
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

	// fmt.Println(listaTareas.tasks[0])

	t3 := &tasks{Nombre: "Completar curso Node Js", Descripcion: "Completar curso en una semana"}
	listaTareas.agregarALista(t3)

	listaTareas.imprimirLista()
	fmt.Println("------------------------------------------------")

	t1.marcarCompleta()
	listaTareas.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Ariel"] = listaTareas

	t4 := &tasks{Nombre: "Completar curso Java", Descripcion: "Completar curso en una semana"}
	t5 := &tasks{Nombre: "Completar curso C#", Descripcion: "Completar curso en una semana"}

	listaTareas2 := &taskList{
		tasks: []*tasks{t4, t5},
	}

	mapaTareas["Ricardo"] = listaTareas2

	fmt.Println("Mapa tareas Ariel")
	mapaTareas["Ariel"].imprimirLista()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("Mapa tareas Ricardo")
	mapaTareas["Ricardo"].imprimirLista()
}
