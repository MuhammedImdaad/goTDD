package main

import (
	"fmt"
	"io"
	// "os"
)

func Greet(w io.Writer, name string) {
	// fmt.Printf("Hello, %s", name) // Printf just calls Fprintf passing in os.Stdout

	// Printf passes os.Stdout to Fprintf which expects an io.Writer
	// os.Stdout implements io.Writer
	fmt.Fprintf(w, "Hello, %s", name)
}

// func main() {
// 	Greet(os.Stdout, "chris")
// }