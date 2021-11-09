package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func IntroHang() {
	//tbHangman := ArrayHangman()
	var choice string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hey you! What is your name ?\n")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Welcome %c!", name)

	fmt.Println("1 - Play")
	fmt.Println("2 - Rules")
	fmt.Println("3 - Level")
	fmt.Println("4 - Leave the game")
	fmt.Println("Press your choice : ")
	fmt.Scanln(&choice)
	if choice == "1" {
		fmt.Println("Let's start the game! ")
		gameHangman()
	} else if choice == "2" {
		fmt.Println()
	}
	Restart()
}

func Restart() {
	var a string
	fmt.Println("if you want to play a HANGMAN GAME? Press X to play, else 0 : ")
	fmt.Scan(&a)
	a = strings.ToUpper(a)
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
	wordHole := WordHoleF(toFind, nbrLShow)
	attempt := 10
	cptFailed := 0
	for attempt > 0 {
		inPut := CheckInPutF(attempt, toFind, wordHole, lTried)
		lTried = append(lTried, inPut)
		if len(inPut) == 1 {
			failed := InPutLettre(inPut, toFind, wordHole)
			if failed == false {
				attempt--
				cptFailed++
				PrintHangman(tbHangman, cptFailed)
			}
		} else {
			if InPutWord(inPut, toFind) == false {
				attempt -= 2
				cptFailed += 2
				PrintHangman(tbHangman, cptFailed)
			}
		}
	}
	if attempt == 0 {
		fmt.Println("YOU LOOSE")
		Restart()
	}
}

func CheckInPutF(attempt int, toFind string, word []rune, lTried []string) string {
	inPut := InPutF(word, attempt)
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
	return inPut
}

func InPutF(word []rune, attempt int) string {
	var inPut string
	fmt.Println(" ")
	fmt.Println(string(word))
	fmt.Printf("You have %v attemps\nChoose : ", attempt)
	fmt.Scan(&inPut)
	inPut = strings.ToLower(inPut)
	return inPut
}

func IsAlpha(s string) bool {
	for _, value := range []rune(s) {
		if (value < 'a' || value > 'z') && value != ' ' {
			fmt.Println("Only letter please")
			return false
		}
	}
	return true
}

func PrintHangman(tbHangman []string, a int) { // fonction qui permet de générer un nouveau nombre random à chaque tour.
	fmt.Println(tbHangman[a])
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

func InPutLettre(inPut, toFind string, wordHole []rune) bool {
	failed := false
	for i := 0; i < len(toFind); i++ {
		if inPut == string(toFind[i]) {
			if wordHole[i] == '_' {
				wordHole[i] = rune(toFind[i])
				failed = true
			}
		}
	}
	if toFind == string(wordHole) {
		fmt.Println(string(wordHole))
		fmt.Printf("Congrats! YOU WON !! \n")
		Restart()
	}
	if failed == true {
		fmt.Println("Well done!")
	}
	return failed
}

func InPutWord(inPut, toFind string) bool {
	a := Compare(inPut, toFind)
	if a == false {
		fmt.Println(toFind)
		fmt.Println("YOU WIN")
		Restart()
	}
	return false
}

func Rules(tbHangman []string, name string) {
	fmt.Print("Hello, "+name, "Time to play hangman!\n")
	time.Sleep(1 * time.Second)
	fmt.Println("Your mission, if you accept it, is to find the secret word.")
	time.Sleep(1 * time.Second)
	fmt.Println("Let us give you some instructions: ")
	time.Sleep(1 * time.Second)
	fmt.Println("We give you a secret word to find.")
	time.Sleep(1 * time.Second)
	fmt.Println("You gonna propose letter or word")
	time.Sleep(1 * time.Second)
	fmt.Println("You gonna loose :\n-1 attempt for each letter failed\n-2 for each word failed")
	fmt.Println("To help you, we gonna reveal some characters" +
		"\nYou have 10 attempts to find it... ")
	time.Sleep(1 * time.Second)
	fmt.Println("If you fail ...")
	time.Sleep(1 * time.Second)
	PrintHangman(tbHangman, 10)
	fmt.Println("José gonna be dead..")
	fmt.Println("Start guessing...")
	time.Sleep(1 * time.Second)

}
