package main

import (
	"bytes"
	"testing"
)
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{} // Buffer type from the bytes package implements the Writer interfac
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}