package hangman

import (
	"fmt"
	"strings"
)

type Game struct {
	WordRoot string
	WordHole []rune
	Attempt  int
	LTried   []string
	InPut    string
}

func Params() Game {
	game := Game{}
	game.WordRoot = ChooseWord()
	game.WordHole = WordHoleF(game.WordRoot)
	game.Attempt = 10

	return game
}

func GameManager(game Game) { // define params to play
	win := GameHangman(game)
	if win == true {
		fmt.Println("You Win!!")
	} else {
		fmt.Println("Looser!!")
	}
	Menu()
}

func Counter(attempt int, win bool) int {
	if win == true && attempt > 0 {
		return attempt
	}
	attempt--
	return attempt
}

func Tested(game Game) ([]string, bool) {
	for _, lettre := range game.LTried { // compare with the letter already tried
		if game.InPut == lettre {
			fmt.Println("you already tried this!")
			return game.LTried, false
		}
	}
	game.LTried = append(game.LTried, game.InPut)
	return game.LTried, true
}

func PrintGame(game Game) {
	fmt.Println(PrintHangman(game.Attempt))
	fmt.Println(string(game.WordHole))
	fmt.Printf("You have %d attempt\n", game.Attempt)
}

func GameHangman(game Game) bool {
	var find, winLettre, tried bool
	for game.Attempt > 0 {
		PrintGame(game)
		game.InPut = InPutF()
		if game.InPut == "--SAVE" {
			Save(game)
			return find
		}
		error := IsAlpha(game.InPut)
		if error == false {
			GameHangman(game)
		}
		game.LTried, tried = Tested(game)
		if tried == false {
			GameHangman(game)
		}
		if len(game.InPut) == 1 { // if it's a letter
			winLettre, find = InPutLettre(game.InPut, game.WordRoot, game.WordHole)
		} else if len(game.InPut) == len(game.WordRoot)-1 {
			find = InPutWord(game.InPut, game.WordRoot)
		} else {
			fmt.Println("Your inPut is not conform, retry please")
			GameHangman(game)
		}
		game.Attempt = Counter(game.Attempt, winLettre)
		if Win(find) == true {
			game.Attempt = 0
		}
	}
	return find
}

func Win(find bool) bool {
	if find == true {
		return true
	}
	return false
}

func InPutF() string { // function to ask and return the input
	var inPut string
	fmt.Scan(&inPut)
	inPut = strings.ToUpper(inPut) // transform the input to Upper Letter
	return inPut                   // return upper input
}

func IsAlpha(s string) bool { // check if the input is only upper alpha character
	for _, value := range []rune(s) {
		if (value >= 'A' && value <= 'Z') || (value >= 'a' && value <= 'z') {
			return true // else return true
		}
	}
	fmt.Println("Only letter please")
	return false // return false if the input is not alpha
}

func InPutLettre(inPut, toFind string, wordHole []rune) (bool, bool) { // if the lengths of the input is 1 so a single letter
	winLettre := false // bool winLettre
	find := false      // bool word find
	for i := 0; i < len(toFind); i++ {
		if inPut == string(toFind[i]) { // if the input is in the word to find
			if wordHole[i] == '_' { // if the letter is hide yet
				wordHole[i] = rune(toFind[i])         // the letter is show
				winLettre = true                      // the bool win letter take the true value
				newWordHole := string(wordHole)       // we update the wordHole
				find = InPutWord(newWordHole, toFind) // we call the function InPutWord, to check is the word is finding. The function return a bol and the result is keeping int the bool Find
			}
		}
	}
	return winLettre, find // the function return the 2 update bool
}

func InPutWord(inPut, toFind string) bool { // the function call the function compare dans return a bool find
	different := Compare(inPut, toFind)
	var find bool
	if different == false {
		fmt.Println(toFind)
		find = true // if there is no difference, inpPut = toFind , find == true
	}
	return find // return update find
}

func Compare(inPut, toFind string) bool { // the function compare index by index the input and the word to find
	different := false
	for indexA := range inPut {
		for i := 0; i < len(toFind); i++ {
			if inPut[indexA] != toFind[indexA] {
				different = true // if there is a difference, different = true
				return different // return true in the function InPutWord
			}
		}
	}
	return different // return update different to function InPutWord
}

func PrintHangman(attempt int) string { // function to print the good hangman step
	HangmanState := ArrayHangman()
	state := 10 - attempt
	return HangmanState[state]
}
