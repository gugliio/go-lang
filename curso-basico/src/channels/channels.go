package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 1)
	fmt.Println("hello")

	go say("Bye", c)

	fmt.Println(<-c)
}
