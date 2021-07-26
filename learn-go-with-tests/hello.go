package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

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
		return spanishHelloPrefix + name
	}
	if lang == "French" {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}
