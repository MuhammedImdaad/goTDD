package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Integration test for recording wins and retrieving player scores
func TestRecordingWinsAndRetrievingThem(t *testing.T) {

	store := NewInMemoryPlayerStore() // Create a new in-memory store
	server := NewPlayerServer(store)  // Create the server with the store
	player := "Pepper"

	// Record three wins for the player
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		// Retrieve the player's score
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		// Assert the score is 3 (three wins recorded)
		assertResponseBody(t, response.Body.String(), "3")
	})
	
	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}
