package concurrency

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup // A WaitGroup waits for a collection of goroutines to finish
		wg.Add(wantedCount)   // to set the number of goroutines to wait for

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done() // inform waitgroup by calling done when finished
			}()
		}

		wg.Wait() //  block until all goroutines have finished

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()

	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
