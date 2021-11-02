package hangman

import (
	"fmt"
)

func IntroHang() {
	Restart()
}

func Restart() {
	var a string
	fmt.Println("Do you want to play a HANGMAN GAME? Press X to play, else 0 : ")
	fmt.Scan(&a)
	if a == "X" {
		gameHangman()
	} else {
		fmt.Println("Oh okay... Hope to see you again, Bye...\n")
		return
	}

}

func gameHangman() {
	toFind := ChooseWord()
	tbHangman := ArrayHangman()
	nbrLShow := len(toFind)/2 - 1

	var lTried []string
	tbWord := PrintWord(toFind, nbrLShow)
	Attempt := 10
	cptFailed := 0
	fmt.Println(string(tbWord))

	for Attempt > 0 {
		failed := false
		guestL := ProposeLetter(Attempt, tbWord, lTried)
		lTried = append(lTried, guestL)
		for i := 0; i < len(toFind); i++ {
			if guestL == string(toFind[i]) {
				tbWord[i] = rune(toFind[i])
				failed = true
			}
		}
		fmt.Println(string(tbWord))
		if toFind == string(tbWord) {
			fmt.Printf("Congrats! YOU WON !! \n")
			Restart()
		}

		if failed == false {
			Attempt--
			cptFailed++
			PrintHangman(tbHangman, cptFailed)
		}
	}
	if Attempt == 0 {
		fmt.Println("YOU LOOSE")
		Restart()
	}
}

func ProposeLetter(Attempt int, tbWord []rune, lTried []string) string {
	var sugestL string
	fmt.Printf("You have %v attemps\nChoose : ", Attempt)
	fmt.Scan(&sugestL)
	sugestL = ToLower(sugestL)

	Err, l := Error(rune(sugestL[0]))

	if Err == true {
		ProposeLetter(Attempt, tbWord, lTried)
	}
	for _, lettre := range lTried {
		if l == rune(lettre[0]) {
			fmt.Println("You already try this letter.\n")
			return ProposeLetter(Attempt, tbWord, lTried)
		}
	}
	return sugestL
}

func ToLower(s string) string {
	var newString string
	for _, lettre := range []rune(s) {
		if lettre >= 'A' && lettre <= 'Z' {
			lettre += 32
		}
		newString += string(lettre)
	}
	return newString
}

func Error(l rune) (bool, rune) {
	if l >= 'a' && l <= 'z' {
		return false, l
	}
	fmt.Printf("Only one letter is accepted, try again\n")
	return true, l
}

func PrintHangman(tbHangman []string, a int) { // fonction qui permet de générer un nouveau nombre random à chaque tour.
	fmt.Println(tbHangman[a])
}
