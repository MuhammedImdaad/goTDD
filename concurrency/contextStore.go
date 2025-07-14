package concurrency

import (
	"fmt"
	"net/http"
)

// Store interface defines methods for fetching data and handling cancellation
// Used to allow the server to interact with any data store that implements these methods
type Store interface {
	Fetch() string
	Cancel()
}

// Server returns an http.HandlerFunc that serves data from the provided Store
// It supports cancellation via request context
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context() // Get the request's context for cancellation support

		data := make(chan string, 1) // Buffered channel to receive fetched data

		// Start fetching data in a separate goroutine
		go func() {
			data <- store.Fetch()
		}()

		// Wait for either the data to be fetched or the context to be cancelled
		select {
		case d := <-data:
			fmt.Fprint(w, d) // Write fetched data to response
		case <-ctx.Done(): // If context is cancelled (client aborts request)
			store.Cancel() // Tell the store to cancel its work
		}
	}
}
