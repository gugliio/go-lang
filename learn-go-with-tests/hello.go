package main

import "fmt"

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return "Hello, world"
}

func HelloName(name string) string {
	return "Hello, " + name
}
