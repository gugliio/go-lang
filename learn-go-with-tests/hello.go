package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return englishHelloPrefix + "world"
}

func HelloName(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}
