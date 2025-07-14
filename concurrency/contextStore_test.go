package concurrency

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store) // returns an http handler function

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request) // calling the http handler function

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store) // returns an http handler function

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// Create a new context (cancellingCtx) derived from the request's context, with a cancel function
		cancellingCtx, cancelFunc := context.WithCancel(request.Context())

		// Schedule the cancel function to be called after 5 milliseconds, simulating a client cancelling the request
		time.AfterFunc(5*time.Millisecond, cancelFunc)

		// Attach the cancelling context to the request
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request) // calling the http handler function

		store.assertWasCancelled()

		/*
			cancellingCtx: a context that will be cancelled after 5ms, simulating a client aborting the request.
			We derive it from the original request context and use it in the request to test cancellation handling.
		*/
	})
}
