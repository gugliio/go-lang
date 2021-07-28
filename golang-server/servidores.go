package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	server := []string{
		"http://platzi.com",
		"http://google.com",
		"http://instagram.com",
		"http://facebook.com",
	}
	for _, server := range server {
		revisarServidor(server)
	}
	tiempoFinal := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoFinal)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println("No esta disponible")
	} else {
		fmt.Println("Esta funcionando normalmente")
	}
}
