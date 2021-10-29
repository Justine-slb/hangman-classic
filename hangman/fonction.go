package hangman

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CacheLettre(toFind string, tbWord []rune, nLshow int) []rune {
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

func PrintHangman(a int) { // fonction qui permet de générer un nouveau nombre random à chaque tour.
	f, e := os.Open("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\hangman.txt")
	check(e)
	defer f.Close()
	data, _ := os.ReadFile("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\hangman.txt")
	tbHangman := strings.SplitN(string(data), "*", -1)
	fmt.Println(tbHangman[a])
}
