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
	channel := make(chan string)
	i := 0
	for {
		if i > 2 {
			break
		}
		for _, server := range server {
			go revisarServidor(server, channel)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)
		i++
	}

	tiempoFinal := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoFinal)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		canal <- servidor + " No esta disponible"
	} else {
		canal <- servidor + " Esta funcionando normalmente"
	}
}
