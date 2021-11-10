package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func IntroHang() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hey you! What is your name ?\n")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Welcome %v\n", name)
	Menu()
}

func Menu() {
	var choice string
	fmt.Println("1 - Play")
	fmt.Println("2 - Rules")
	fmt.Println("3 - Level")
	fmt.Println("4 - Leave the game")
	fmt.Println("Press your choice : ")
	fmt.Scanln(&choice)
	if choice == "1" {
		Params()
	} else if choice == "2" {
		Rules()
	} else if choice == "3" {
		Params()
		//HangmanAsciiArt()
	} else if choice == "4" {
		fmt.Println("See you soon!")
		return
	}
}

func Rules() {
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
	Menu()
}

func Restart() {
	var a string
	CallClear()
	fmt.Println("If you want restart, press X else press 0 ")
	fmt.Scan(&a)
	a = strings.ToUpper(a)
	if a == "X" {
		Params()
	} else {
		Menu()
	}
}
