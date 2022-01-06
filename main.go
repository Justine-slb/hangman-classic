package main

import (
	"hangman/hangman"
	"os"
)

func main() {

	if len(os.Args) == 3 {
		hangman.Load()
	}
	hangman.IntroHang()
}
