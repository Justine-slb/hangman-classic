package main

import (
	"fonction.go/hangman"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		hangman.Load()
	}
	hangman.Menu()
}
