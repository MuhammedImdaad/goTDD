package main

import (
	"log"
	"net/http"
	"os"
)

// Entry point for the application
// Starts the HTTP server on port 5000 using PlayerServer
func main() {
	const dbFileName = "game.db.json"
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := NewPlayerServer(store)
	// Listen on port 5000 and serve requests using PlayerServer
	log.Fatal(http.ListenAndServe(":5000", server))

	/*
		POST - curl -X POST http://localhost:5000/players/{player name}
		GET - curl http://localhost:5000/players/{player name}
	*/
}
