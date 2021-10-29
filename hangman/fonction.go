package hangman

import (
	"fmt"
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
func ProposeLetter(Attempt int, tbWord []rune, sugestL string, lTried []string) string {
	fmt.Printf("You have %v attemps\n", Attempt)
	fmt.Println(string(tbWord))
	fmt.Printf("Propose a lower letter\n")
	fmt.Scan(&sugestL)
	if Error(sugestL) == true {
		ProposeLetter(Attempt, tbWord, sugestL, lTried)
	}
	for _, value := range lTried {
		if rune(sugestL[0]) == rune(value[0]) {
			fmt.Println("You already try this letter, try an other")
			ProposeLetter(Attempt, tbWord, sugestL, lTried)
		}
	}
	return sugestL
}

func Error(l string) bool {
	if rune(l[0]) < 97 || rune(l[0]) > 122 || len(l) > 1 {
		fmt.Printf("Only one lower letter is accepted, try again\n")
		return true
	}
	return false
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

func GoodLetter(i int, tbWord []rune, toFind string) bool {
	fmt.Println("Well done!\n")
	tbWord[i] = rune(toFind[i])
	if string(tbWord) != toFind {
		return false
	} else {
		fmt.Println(string(tbWord))
		fmt.Printf("You won\n")
		return true
	}
}

func PrintHangman(tbHangman []string, a int) { // fonction qui permet de générer un nouveau nombre random à chaque tour.
	fmt.Println(tbHangman[a])
}

func ArrayHangman() []string {
	f, e := os.Open("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\hangman.txt")
	check(e)
	defer f.Close()
	data, _ := os.ReadFile("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\hangman.txt")
	tbHangman := strings.SplitN(string(data), "*", -1)
	return tbHangman
}
