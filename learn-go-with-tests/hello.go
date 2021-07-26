package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return englishHelloPrefix + "world"
}

func HelloName(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
