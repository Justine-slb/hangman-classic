package hangman

import (
	"fmt"
	"strings"
	"time"
)

func Params(asciiOk bool) { // define params to play
	toFind := ChooseWord()                                  // the hide word to find
	nbrLShow := len(toFind)/2 - 1                           // number of letter to show at the beginning of the game
	var lTried []string                                     // init a variable to keep the letter and word already tried
	wordHole := WordHoleF(toFind, nbrLShow)                 // word show to begin the game
	attempt := 10                                           // number of attempt for the player
	GameHangman(wordHole, toFind, attempt, lTried, asciiOk) // function GameHangman call with all params
}

func GameHangman(wordHole []rune, toFind string, attempt int, lTried []string, asciiOk bool) {
	tbHangman := ArrayHangman() // Array with all the steps to the hangman, they will be print one by one when the player loose attempts
	find := false
	winLettre := false
	cptFailed := 10 // number of failed try. This variable will use to print the good hangman step

	for attempt > 0 { // first for loop, if attempt = 0; the game leave the loop and the game is finish.
		CallClear()                                                      // unload the screen
		cptFailed = 10 - attempt                                         // upload the number of failed attempt
		PrintHangman(tbHangman, cptFailed)                               // Print hangman step
		inPut := CheckInPutF(attempt, toFind, wordHole, lTried, asciiOk) // call the function to chek if the in put is compliant to the rules and to check
		if inPut == "--STOP" {                                           //
			Save(wordHole, toFind, attempt, lTried, asciiOk) // call the save function to save the data and stop the game
			return

		} else {
			lTried = append(lTried, inPut) // if the input is ok, you put it in the array ltried
			if len(inPut) == 1 {           // if it's a letter
				winLettre, find = InPutLettre(inPut, toFind, wordHole) // call the function InPutLettre and return 2 booleans one for the word and one for the letter
				if winLettre == false {
					attempt-- // loose 1 attempt
				} else if find == true {
					attempt = 0 // init attempt to end the game and recall the menu
				}
			} else { // if the inPut is a word
				if InPutWord(inPut, toFind) == false { // return a boolean, if failed :
					attempt -= 2 // loose two attempt
				} else { // else if InPutWord == true
					attempt = 0 // init attempt to 0 to end the run and recall the menu
				}
			}
		}
	}
	if attempt == 0 && cptFailed >= 10 { // if cptFailed >=10
		fmt.Println("YOU LOOSE") // you loose
	}
	Menu(asciiOk) // recall the menu
}

func CheckInPutF(attempt int, toFind string, word []rune, lTried []string, asciiOk bool) string { // check the input and return it when is compliant to the rules
	inPut := InPutF(word, attempt, asciiOk) // ask the input
	if inPut == "--STOP" {                  // command to stop the game and return to the HangmanGame function
		return inPut
	} else {
		err := IsAlpha(inPut) // check if the input is alpha character
		if err == false {     // if err == false recall the function to ask a new input
			return CheckInPutF(attempt, toFind, word, lTried, asciiOk)
		} else if len(inPut) != len(toFind)-1 && len(inPut) != 1 { // else if it is alpha input but not a single letter and not a word of the size of the word to find
			a := len(toFind)
			fmt.Printf("The number of argument is not exact. You can try a letter or a word with %v arguments.\n\n", a)
			return CheckInPutF(attempt, toFind, word, lTried, asciiOk) // recall the function to ask a new intput
		} // else continue, the input is alpha, with a single letter or a word same size to the word to find
		for _, lettre := range lTried { // compare with the letter already tried
			if inPut == lettre {
				fmt.Println("You already try this.\n")
				return CheckInPutF(attempt, toFind, word, lTried, asciiOk) // if the letter is already tried, recall the function
			}
		}
	}
	return inPut // when everything is ok, return the good input to test in the HangmanGame function
}

func InPutF(word []rune, attempt int, asciiOk bool) string { // function to ask and return the input
	var inPut string
	fmt.Println(" ")
	fmt.Println(asciiOk)
	if asciiOk == true {
		AsciiArt(string(word))
	} else {
		fmt.Println(string(word))
	}

	fmt.Printf("You have %v attemps\nChoose : ", attempt)
	fmt.Scan(&inPut)
	inPut = strings.ToUpper(inPut) // transform the input to Upper Letter
	return inPut                   // return upper input
}

func IsAlpha(s string) bool { // check if the input is only upper alpha character
	for _, value := range []rune(s) {
		if value < 'A' || value > 'Z' {
			fmt.Println("Only letter please")
			return false // return false if the input is not alpha
		}
	}
	return true // else return true
}

func InPutLettre(inPut, toFind string, wordHole []rune) (bool, bool) { // if the lengths of the input is 1 so a single letter
	winLettre := false // bool winLettre
	find := false      // bool word find
	for i := 0; i < len(toFind); i++ {
		if inPut == string(toFind[i]) { // if the input is in the word to find
			if wordHole[i] == '_' { // if the letter is hide yet
				wordHole[i] = rune(toFind[i]) // the letter is show
				fmt.Println("Well done!")
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
		fmt.Println("YOU WIN")
		time.Sleep(2 * time.Second)
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

func PrintHangman(tbHangman []string, a int) { // function to print the good hangman step
	fmt.Println(tbHangman[a])
}
