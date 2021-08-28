package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.6888, -74.25698}
	fmt.Println(m["Bell Labs"])

	m2 := map[string]Vertex{
		"Vertex1": Vertex{40.1234, -74.25698},
		"Vertex2": Vertex{40.4321, -74.2569888},
	}
	fmt.Println(m2)

	m3 := map[string]Vertex{
		"Vertex1": {40.1234, -74.25698},
		"Vertex2": {40.4321, -74.2569888},
	}
	fmt.Println(m3)

	m4 := make(map[string]int)
	m4["Test"] = 42
	fmt.Println(m4["Test"])

	m4["Test"] = 50
	fmt.Println(m4["Test"])
}
