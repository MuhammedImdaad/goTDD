package server

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileSystemPlayerStore struct { // implement PlayerStore
	database *os.File
	league   League
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}

	return nil
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file, %v", err)
	}
	return &FileSystemPlayerStore{
		database: file,
		league:   league,
	}, nil
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})

	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	player := f.league.Find(name)
	if player != nil {
		wins = player.Wins
	}
	return wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++ //player pointer
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(f.league)
}
