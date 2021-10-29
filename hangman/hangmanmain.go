package hangman

import (
	"fmt"
)

func HangmanMain() {
	toFind := ChooseWord()
	tbHangman := ArrayHangman()
	nbrLShow := len(toFind)/2 - 1
	var sugestL string
	tbWord := []rune(toFind)
	for n := 0; n < len(tbWord)-1; n++ {
		tbWord[n] = '_'
	}

	var failed bool
	var lTried []string

	tbWord = CacheLettre(toFind, tbWord, nbrLShow)

	Attempt := 10
	cptFailed := 0

	for Attempt > 0 {
		if Attempt != 10 {
			fmt.Printf("You already try this letters : %v \n", lTried)
		}
		sugestL = ProposeLetter(Attempt, tbWord, sugestL, lTried)
		lTried = append(lTried, sugestL)
		for i := 0; i < len(toFind); i++ {
			if sugestL == string(toFind[i]) {
				failed = true
				if GoodLetter(i, tbWord, toFind) == true {
					return
				}
			}
		}
		if failed == false {
			Attempt--
			cptFailed++
			PrintHangman(tbHangman, cptFailed)
		}
		failed = false
	}
	if Attempt == 0 {
		fmt.Println("you loose")
	}
}
