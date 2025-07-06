package main // Programs start running in package main

import (
	"fmt"
	"rsc.io/quote" // `go mod tidy` will add the quote module as a requirement
)

const englishHelloPrefix = "Hello, " // Constants cannot be declared using the := syntax

func Hello(name string) string { // type comes after the variable name.
	if name == "" {
		return englishHelloPrefix + "World"
	}

	return englishHelloPrefix + name
}

func main() { // main function, go run .
	fmt.Println(Hello("imd"))
	fmt.Println(quote.Go())
}