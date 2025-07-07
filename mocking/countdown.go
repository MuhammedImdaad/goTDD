package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface{
	Sleep()
}

type RealSleeper struct{}

func (r *RealSleeper) Sleep() {
	time.Sleep(time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout, &RealSleeper{})
}
