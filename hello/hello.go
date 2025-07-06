package main // Programs start running in package main

import (
	"fmt"
	"rsc.io/quote" // `go mod tidy` will add the quote module as a requirement
)

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, " // Constants cannot be declared using the := syntax
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string { // type comes after the variable name.
	if name == "" {
		name = "World"
	}

	switch language {
	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() { // main function, go run .
	fmt.Println(Hello("imd", ""))
	fmt.Println(quote.Go())
}
