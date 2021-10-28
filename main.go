package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func SplitWhiteSpaces(s string) []string {
	var sliceFinal []string
	var mot string

	for _, value := range s {
		if value != '\n' && value != 0 {
			mot += string(value)
		} else {
			if mot != "" {
				sliceFinal = append(sliceFinal, mot)
				mot = ""
			}
		}
	}
	if mot != "" {
		sliceFinal = append(sliceFinal, mot)
	}
	return sliceFinal
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	data, _ := os.ReadFile("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\words2.txt")
	fmt.Println(string(data))
	tbWord := SplitWhiteSpaces(string(data))
	ToFind := tbWord[rand.Intn(len(tbWord))]
	Attempt := 10
	var word string

}
