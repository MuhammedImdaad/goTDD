package main

import (
	"github.com/MuhammedImdaad/goTDD/server"
	"log"
	"net/http"
)

// Entry point for the application
// Starts the HTTP server on port 5000 using PlayerServer
func main() {
	const dbFileName = "game.db.json"
	store, close, err := server.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := server.NewPlayerServer(store)
	// Listen on port 5000 and serve requests using PlayerServer
	log.Fatal(http.ListenAndServe(":5000", server))

	/*
		POST - curl -X POST http://localhost:5000/players/{player name}
		GET - curl http://localhost:5000/players/{player name}
	*/
}
