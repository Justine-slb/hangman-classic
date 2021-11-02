package hangman

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

func ChooseWord() string {
	rand.Seed(time.Now().UnixNano())
	f, e := os.Open("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\dictionnaire.txt")
	check(e)
	defer f.Close()
	data, _ := os.ReadFile("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\dictionnaire.txt")
	tbWords := strings.SplitN(string(data), "\n", -1)
	toFind := tbWords[rand.Intn(len(tbWords))]
	return toFind
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ArrayHangman() []string {
	f, e := os.Open("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\hangman.txt")
	check(e)
	defer f.Close()
	data, _ := os.ReadFile("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\hangman.txt")
	tbHangman := strings.SplitN(string(data), "*", -1)
	return tbHangman
}

func PrintWord(toFind string, nLshow int) []rune {
	tbWord := []rune(toFind)
	for n := 0; n < len(tbWord)-1; n++ {
		tbWord[n] = '_'
	}
	tb := HideLetters(nLshow, tbWord)
	for i := 0; i < len(tb); i++ {
		for j := 0; j < len(toFind); j++ {
			if j == tb[i] {
				tbWord[tb[i]] = rune(toFind[tb[i]])
			}
		}
	}
	return tbWord
}

func HideLetters(nLToShow int, tbWord []rune) []int {
	rand.Seed(time.Now().UnixNano())
	var tb []int
	for n := 0; n < len(tbWord)-1; n++ {
		tbWord[n] = '_'
	}
	for i := 0; i < nLToShow; i++ {
		Rnbr := rand.Intn(len(tbWord) - 1)
		for j := 0; j < len(tb); j++ {
			for Rnbr == tb[j] {
				return HideLetters(nLToShow, tbWord)
			}
		}
		tb = append(tb, Rnbr)
	}
	return tb
}
