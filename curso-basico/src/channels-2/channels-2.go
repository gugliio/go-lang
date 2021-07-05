package main

import "fmt"

func message(s string, c chan string) {
	c <- s
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println(len(c), cap(c))

	// Close (cerrar canal una vez que se deja de usuar)
	close(c)
	//c <- "Mensaje 3"

	// Range
	for message := range c {
		fmt.Println(message)
	}

	//Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1", email1)
	go message("mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Mensaje recibido de Email 1", m1)
		case m2 := <-email2:
			fmt.Println("Mensaje recibido de Email 2", m2)
		}
	}
}
