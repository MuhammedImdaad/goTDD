package main

import "testing"

/*
test files needs to be in a file with a name like xxx_test.go
the test function must start with the word Test
*/

func TestHello(t *testing.T) { // go run test
	t.Run("saying hello to people", func(t *testing.T) { // subtests
		got := Hello("imd", "")
		var want string = "Hello, imd"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French for the world", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) { // testing.TB is an interface that *testing.T and *testing.B both satisfy
	t.Helper() // this method is a helper, when it fails, the line number reported will be in our function call rather than inside our test helper
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
