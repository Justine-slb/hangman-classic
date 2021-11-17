package main

import (
	"fmt"
	"fonction.go/hangman"
	"os"
)

func main() {
	var asciiOk bool
	if len(os.Args) == 4 {
		fmt.Println("+")
		asciiOk = hangman.AsciiArtok()
	}
	if len(os.Args) == 3 {

		hangman.Load()
	}
	hangman.IntroHang(asciiOk)
}
