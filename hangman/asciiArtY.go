package hangman

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func AsciiArtok() bool {
	if os.Args[1] == "words.txt" {
		if os.Args[2] == "--letterFile" {
			if os.Args[3] == "standard.txt" {
				return true
			}
		}
	}
	return false
}

func AsciiArt(word string) {
	f, e := os.Open("standard.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	data, err := os.ReadFile("standard.txt")

	if err != nil {
		log.Fatal(err)
	}
	Asciiart := strings.Split(string(data), "\r\n")
	R := []rune(word)

	for i := 0; i <= len(word)-1; i++ {
		if word[i] >= ' ' && word[i] <= '~' {
			R[i] = R[i] - 32
			index := int(R[i])
			fmt.Print(Asciiart[index*9+1])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(word)-1; i++ {
		if word[i] >= ' ' && word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+2])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(word)-1; i++ {
		if word[i] >= ' ' && word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+3])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(word)-1; i++ {
		if word[i] >= ' ' && word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+4])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(word)-1; i++ {
		if word[i] >= ' ' && word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+5])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(word)-1; i++ {
		if word[i] >= ' ' && word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+6])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(word)-1; i++ {
		if word[i] >= ' ' && word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+7])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(word)-1; i++ {
		if word[i] >= ' ' && word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+8])
		}
	}
	fmt.Print("\n\n")
}
