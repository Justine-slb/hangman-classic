package hangman

import (
	"fmt"
	"time"
)

func Menu() { // function to print the menu
	var choice string
	CallClear()
	word := "HANGMAN"
	AsciiArt(word)
	fmt.Println("1 - Play")
	fmt.Println("2 - Rules")
	fmt.Println("3 - Game started")
	fmt.Println("4 - Leave")
	fmt.Println("Press your choice : ")
	fmt.Scan(&choice)
	if choice == "1" {
		game := Init(Level())
		GameManager(game)
	} else if choice == "2" {
		Rules()
	} else if choice == "3" {
		game := Load()
		GameManager(game)
	} else if choice == "4" {
		fmt.Println("See you soon!")
		return
	} else {
		fmt.Println("Input not recognized, try again:")
		time.Sleep(2 * time.Second)
		Menu()
	}
}

func Rules() { // function to print the Rules
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
	PrintHangman(0)
	time.Sleep(3 * time.Second)
	fmt.Println("If you want to stop the game and save your run, you can input '--stop' in the terminal")
	Menu()
}
