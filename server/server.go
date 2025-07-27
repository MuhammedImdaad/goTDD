package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore interface defines methods for getting a player's score and recording a win
// Used by PlayerServer to interact with the underlying data store
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

// PlayerServer handles HTTP requests for player scores and win recording
// It uses a PlayerStore to manage player data

type PlayerServer struct {
	// implement http Handler interface via ServeHTTP method
	// ServeHTTP method exposed via the embedded handler

	store        PlayerStore
	http.Handler //  Go's lightweight alternative to inheritance
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	// built-in routing mechanism (request multiplexer) which lets you attach http.Handlers to particular request paths
	// ServeHTTP routes incoming HTTP requests to the appropriate handler based on method and path

	r := http.NewServeMux() // ServeMux is also an http.Handler
	r.Handle("/league", http.HandlerFunc(p.leagueHandler))
	r.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.Handler = r

	return p
}

const jsonContentType = "application/json"
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
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
