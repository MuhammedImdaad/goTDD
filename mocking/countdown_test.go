package main

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	sleep = "s"
	write = "p"
)

// implements both io.Writer and Sleeper, recording every call into one slice
type SpyCountdownOperations struct {
	S []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.S = append(s.S, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.S = append(s.S, write)
	return
}

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.S) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.S)
		}
	})
}
