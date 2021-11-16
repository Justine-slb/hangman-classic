package hangman

import (
	"fmt"
	"strings"
	"time"
)

func Params() {
	toFind := ChooseWord()
	nbrLShow := len(toFind)/2 - 1
	var lTried []string
	wordHole := WordHoleF(toFind, nbrLShow)
	attempt := 10
	GameHangman(wordHole, toFind, attempt, lTried)
}

func GameHangman(wordHole []rune, toFind string, attempt int, lTried []string) {
	tbHangman := ArrayHangman()
	find := false
	winLettre := false
	cptFailed := 10

	for attempt > 0 {
		CallClear()
		cptFailed = 10 - attempt
		PrintHangman(tbHangman, cptFailed)
		inPut := CheckInPutF(attempt, toFind, wordHole, lTried)
		if inPut == "STOP" {
			Save(wordHole, toFind, attempt, lTried)
			return

		} else {
			lTried = append(lTried, inPut)
			if len(inPut) == 1 {
				winLettre, find = InPutLettre(inPut, toFind, wordHole)
				if winLettre == false {
					attempt--
				} else if find == true {
					attempt = 0
				}
			} else {
				if InPutWord(inPut, toFind) == false {
					attempt -= 2
					cptFailed += 2
				} else {
					attempt = 0
				}
			}
		}
	}
	if attempt == 0 && cptFailed >= 10 {
		fmt.Println("YOU LOOSE")
	}
	Menu()
}

func CheckInPutF(attempt int, toFind string, word []rune, lTried []string) string {
	inPut := InPutF(word, attempt)
	if inPut == "STOP" {
		return inPut
	} else {
		err := IsAlpha(inPut)
		if err == false {
			return CheckInPutF(attempt, toFind, word, lTried)
		} else if len(inPut) != len(toFind)-1 && len(inPut) != 1 {
			a := len(toFind)
			fmt.Printf("The number of argument is not exact. You can try a letter or a word with %v arguments.\n\n", a)
			return CheckInPutF(attempt, toFind, word, lTried)
		}
		for _, lettre := range lTried {
			if inPut == lettre {
				fmt.Println("You already try this.\n")
				return CheckInPutF(attempt, toFind, word, lTried)
			}
		}
	}
	return inPut
}

func InPutF(word []rune, attempt int) string {
	var inPut string
	fmt.Println(" ")
	fmt.Println(string(word))
	fmt.Printf("You have %v attemps\nChoose : ", attempt)
	fmt.Scan(&inPut)
	inPut = strings.ToUpper(inPut)
	return inPut
}

func IsAlpha(s string) bool {
	for _, value := range []rune(s) {
		if (value < 'A' || value > 'Z') && value != ' ' {
			fmt.Println("Only letter please")
			return false
		}
	}
	return true
}

func InPutLettre(inPut, toFind string, wordHole []rune) (bool, bool) {
	winLettre := false
	find := false

	for i := 0; i < len(toFind); i++ {
		if inPut == string(toFind[i]) {
			if wordHole[i] == '_' {
				wordHole[i] = rune(toFind[i])
				fmt.Println("Well done!")
				winLettre = true
				newWordHole := string(wordHole)
				find = InPutWord(newWordHole, toFind)
			}
		}
	}
	return winLettre, find
}

func InPutWord(inPut, toFind string) bool {
	different := Compare(inPut, toFind)
	var find bool
	if different == false {
		fmt.Println(toFind)
		fmt.Println("YOU WIN")
		time.Sleep(2 * time.Second)
		find = true
	}
	return find
}

func Compare(inPut, toFind string) bool {
	different := false
	for indexA := range inPut {
		for i := 0; i < len(toFind); i++ {
			if inPut[indexA] != toFind[indexA] {
				different = true
				return different
			}
		}
	}
	return different
}

func PrintHangman(tbHangman []string, a int) {
	fmt.Println(tbHangman[a])
}
