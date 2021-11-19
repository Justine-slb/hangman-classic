package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func IntroHang(asciiOk bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hey you!\nWelcome to the HANGMAN GAME\nWhat is your name ?\n")
	name, _ := reader.ReadString('\n')
	name = strings.ToUpper(name)
	fmt.Println("HELLO" + name)
	fmt.Println("TIME TO PLAY HANGMAN!")
	time.Sleep(2 * time.Second)
	CallClear()
	Menu(asciiOk)
}

func Menu(asciiOk bool) { // function to print the menu
	var choice string
	CallClear()
	word := "HANGMAN"
	AsciiArt(word)
	fmt.Println("1 - Play")
	fmt.Println("2 - Rules")
	fmt.Println("3 - Ascii Art mode")
	fmt.Println("4 - Leave")
	fmt.Println("Press your choice : ")
	fmt.Scan(&choice)
	if choice == "1" {
		Params(asciiOk)
	} else if choice == "2" {
		Rules(asciiOk)
	} else if choice == "3" {
		asciiOk = true
		Params(asciiOk)
	} else if choice == "4" {
		fmt.Println("See you soon!")
		return
	} else {
		fmt.Println("Input not recognized, try again:")
		time.Sleep(2 * time.Second)
		Menu(asciiOk)
	}
}

func Rules(asciiOk bool) { // function to print the Rules
	tbHangman := ArrayHangman()
	fmt.Println("Your mission, if you accept it, is to find the secret word.")
	time.Sleep(1 * time.Second)
	fmt.Println("Let us give you some instructions: ")
	time.Sleep(1 * time.Second)
	fmt.Println("We give you a secret word to find.")
	time.Sleep(1 * time.Second)
	fmt.Println("You gonna propose letter or word")
	time.Sleep(1 * time.Second)
	fmt.Println("You gonna loose :\n-1 attempt for each letter failed")
	time.Sleep(1 * time.Second)
	fmt.Println("-2 for each word failed")
	time.Sleep(1 * time.Second)
	fmt.Println("To help you, we gonna reveal some characters" +
		"\nYou have 10 attempts to find it... ")
	time.Sleep(1 * time.Second)
	fmt.Println("If you fail ...")
	time.Sleep(1 * time.Second)
	PrintHangman(tbHangman, 10)
	fmt.Println()
	fmt.Println("If you want to stop the game and save your run, you can in put '--stop' in the terminal")
	Menu(asciiOk)
}
