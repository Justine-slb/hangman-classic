package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func IntroHang() {
	tbHangman := ArrayHangman()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hey you! What is your name ?\n")
	name, _ := reader.ReadString('\n')
	fmt.Print("Hello, "+name, "Time to play hangman!\n")
	time.Sleep(2 * time.Second)
	fmt.Println("Your mission, if you accept it, is too find the secret word.")
	time.Sleep(2 * time.Second)
	fmt.Println("Let us give you some instructions: ")
	time.Sleep(2 * time.Second)
	fmt.Println("We give you a secret word to find.")
	time.Sleep(2 * time.Second)
	fmt.Println("You gonna propose letter or word")
	time.Sleep(2 * time.Second)
	fmt.Println("You gonna loose 1 attempt for each letter failed and 2 for each word failed")
	fmt.Println("To help you, we gonna reveal some caractere\nYou have 10 attempts to find it... ")
	time.Sleep(2 * time.Second)
	fmt.Println("If you fail ...")
	time.Sleep(2 * time.Second)
	PrintHangman(tbHangman, 10)
	fmt.Println("José gonna be dead..")
	fmt.Println("Start guessing...")
	time.Sleep(1 * time.Second)
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
	word := PrintWord(toFind, nbrLShow)
	attempt := 10
	cptFailed := 0

	for attempt > 0 {
		failed := false
		guestL := ProposeLetter(attempt, toFind, word, lTried)
		lTried = append(lTried, guestL)
		if len(guestL) == 1 {
			for i := 0; i < len(toFind); i++ {
				if guestL == string(toFind[i]) {
					if word[i] == '_' {
						word[i] = rune(toFind[i])
						failed = true
					}
				}
			}
			fmt.Println(string(word))
			if toFind == string(word) {
				fmt.Printf("Congrats! YOU WON !! \n")
				Restart()
			}
			if failed == false {
				attempt--
				cptFailed++
				PrintHangman(tbHangman, cptFailed)
			}
		} else {
			a := Compare(guestL, toFind)
			fmt.Println(a)
			if a == false {
				fmt.Println("YOU WIN")
				Restart()
			} else {
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

func ProposeLetter(attempt int, toFind string, word []rune, lTried []string) string {
	var sugestL string
	fmt.Println(" ")
	fmt.Println(string(word))
	fmt.Printf("You have %v attemps\nChoose : ", attempt)
	fmt.Scan(&sugestL)
	sugestL = ToLower(sugestL)

	if len(sugestL) == 1 {
		err := IsAlpha(sugestL)
		if err == false {
			fmt.Println("You can only propose letter or word, special caracteres and numbers are not accepted\n")
			ProposeLetter(attempt, toFind, word, lTried)
		}
	} else if len(sugestL) == len(toFind)-1 {
		err := IsAlpha(sugestL)
		for i := 0; i < len(sugestL); i++ {
			if err == false {
				ProposeLetter(attempt, toFind, word, lTried)
			}
		}
	} else if len(sugestL)+1 != len(toFind) {
		a := len(toFind)
		fmt.Printf("The number of argument is not exact. You can try a letter or a word with %v arguments.\n\n", a)
		ProposeLetter(attempt, toFind, word, lTried)
	}

	for _, lettre := range lTried {
		if sugestL == lettre {
			fmt.Println("You already try this.\n")
			return ProposeLetter(attempt, toFind, word, lTried)
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

func Compare(a, b string) bool {
	different := false
	for indexA := range a {
		for i := 0; i < len(b); i++ {
			if a[indexA] != b[indexA] {
				different = true
				return different
			}
		}
	}
	return different
}
