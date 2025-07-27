package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// StubPlayerStore is an implementation of the PlayerStore interface for testing purposes.
type StubPlayerStore struct { // implement PlayerStore interface
	scores   map[string]int // stores player scores
	winCalls []string       // records calls to RecordWin
}

// GetPlayerScore returns the score of a player.
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

// RecordWin records a win for a player.
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// TestGETPlayers tests the PlayerServer GET /players/{name} and POST /players/{name} endpoints.
func TestGETPlayers(t *testing.T) {
	// Create a stub store with initial scores and no win calls
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		winCalls: nil,
	}

	// Create the server using the stub store
	server := &PlayerServer{&store}

	// Test: returns Pepper's score
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	// Test: returns Floyd's score
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	// Test: returns 404 on missing players
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})

	// Test: returns accepted on POST
	t.Run("it returns accepted on POST", func(t *testing.T) {
		request := newPostWinRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})

	// Test: records wins when POST
	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"
		store.winCalls = nil // Reset winCalls before test
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		// Check that RecordWin was called exactly once
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		// Check that the correct player was recorded as winner
		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

// Helper to create a GET request for a player's score
func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// Helper to create a POST request to record a win for a player
func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// Helper to assert response body matches expectation
func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

// Helper to assert HTTP status code matches expectation
func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
