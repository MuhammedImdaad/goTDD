package main

import (
	"fmt"
	"github.com/MuhammedImdaad/goTDD/server"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := server.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := server.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
