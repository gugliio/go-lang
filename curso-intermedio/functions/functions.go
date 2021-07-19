package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5

	y := func() int {
		return x * 2
	}()
	fmt.Println(y)

	/* z := 10
	w := func() int {
		return z * 2
	}()
	fmt.Println(w)*/
	c := make(chan int)
	go func() {
		fmt.Println("Inicio de funcion")
		time.Sleep(5 * time.Millisecond)
		c <- 1
	}()
	<-c
}
