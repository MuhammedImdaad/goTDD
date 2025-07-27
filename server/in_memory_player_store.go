package main

// NewInMemoryPlayerStore creates a new in-memory store for player scores
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore implements PlayerStore using a map for storage
type InMemoryPlayerStore struct {
	store map[string]int // maps player names to their scores
}

// RecordWin increments the win count for the given player
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetPlayerScore returns the score for the given player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
