package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//Recorrer map
	for key, value := range m {
		fmt.Println(key, value)
	}
	//Encontrar un valor
	value := m["Jose"]
	fmt.Println(value)
}
