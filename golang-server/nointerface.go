package main

import "fmt"

type perro struct {
}

type pez struct {
}

type pajaro struct {
}

func (perro) caminar() string {
	return "soy un perro y camino"
}

func (pez) nadar() string {
	return "soy un pez y nado"
}

func (pajaro) volar() string {
	return "soy un pajaro y estoy volando"
}

func moverPerro(perro perro) {
	fmt.Println(perro.caminar())
}

func moverPez(pez pez) {
	fmt.Println(pez.nadar())
}
func moverPajaro(pajaro pajaro) {
	fmt.Println(pajaro.volar())
}

func main() {
	p := perro{}
	moverPerro(p)
	pe := pez{}
	moverPez(pe)
	pa := pajaro{}
	moverPajaro(pa)
}
