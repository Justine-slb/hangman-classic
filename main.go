package main

import (
	"hangman/hangman"
	"os"
)

type Game struct {
	WordRoot string
	WordHole []rune
	Attempt  int
	lTried   []string
	InPut    string
}

func main() {

	if len(os.Args) == 3 {
		hangman.Load()
	}
	hangman.Menu()
}
