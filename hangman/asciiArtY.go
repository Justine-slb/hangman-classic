package hangman

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func AsciiArt() {
	Word := "deeznut"
	f, e := os.Open("C:\\Users\\Nguye\\Downloads\\standard (2).txt")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	data, err := os.ReadFile("C:\\Users\\Nguye\\Downloads\\standard (2).txt")

	if err != nil {
		log.Fatal(err)
	}
	Asciiart := strings.Split(string(data), "\r\n")
	R := []rune(Word)

	for i := 0; i <= len(Word)-1; i++ {
		if Word[i] >= ' ' && Word[i] <= '~' {
			R[i] = R[i] - 32
			index := int(R[i])
			fmt.Print(Asciiart[index*9+1])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(Word)-1; i++ {
		if Word[i] >= ' ' && Word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+2])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(Word)-1; i++ {
		if Word[i] >= ' ' && Word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+3])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(Word)-1; i++ {
		if Word[i] >= ' ' && Word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+4])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(Word)-1; i++ {
		if Word[i] >= ' ' && Word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+5])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(Word)-1; i++ {
		if Word[i] >= ' ' && Word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+6])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(Word)-1; i++ {
		if Word[i] >= ' ' && Word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+7])
		}
	}
	fmt.Print("\n")
	for i := 0; i <= len(Word)-1; i++ {
		if Word[i] >= ' ' && Word[i] <= '~' {
			R[i] = R[i]
			index := int(R[i])
			fmt.Print(Asciiart[index*9+8])
		}
	}
}
