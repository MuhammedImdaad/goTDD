package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore interface defines methods for getting a player's score and recording a win
// Used by PlayerServer to interact with the underlying data store
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer handles HTTP requests for player scores and win recording
// It uses a PlayerStore to manage player data
type PlayerServer struct { // implement http Handler interface via ServeHTTP method
	store PlayerStore
}

// ServeHTTP routes incoming HTTP requests to the appropriate handler based on method and path
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player) // Handle POST: record a win
	case http.MethodGet:
		p.showScore(w, player) // Handle GET: show player's score
	}
}

// showScore writes the player's score to the response
// Returns 404 if the player is not found
func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

// processWin records a win for the player and returns HTTP 202 Accepted
func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	w.WriteHeader(http.StatusAccepted)
	p.store.RecordWin(player)
}
