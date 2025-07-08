package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

type empty struct{} // chan struct{} is the smallest data type available

const timeOut = time.Second * 10

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, timeOut)
}

func ConfigurableRacer(a, b string, tm time.Duration) (string, error) {
	// select lets a goroutine wait on multiple communication operations
	// include time.After in one of your cases to prevent your system blocking forever
	timeOut := time.After(tm)
	select { // The first one to send a value "wins"
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-timeOut:
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
func ping(url string) chan empty {
	ch := make(chan empty) // Always make channels
	go func() {
		http.Get(url)
		close(ch) // close a channel to indicate that no more values will be sent
		// Closing is only necessary when the receiver must be told there are no more values coming
	}()
	return ch
}
