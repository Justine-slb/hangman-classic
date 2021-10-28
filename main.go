package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type HangManData struct {
	Word             string     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
}

func Error(l string) {
	if rune(l[0]) < 97 || rune(l[0]) > 122 || len(l) > 1 {
		fmt.Printf("only one lower letter is accepted, try again\n")
		fmt.Printf("Propose a new lower letter\n")
		fmt.Scan(&l)
		Error(l)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // fonction qui permet de générer un nouveau nombre random à chaque tour.
	f, e := os.Open("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\words2.txt")
	check(e)
	defer f.Close()
	data, _ := os.ReadFile("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\words2.txt")

	tbWord := strings.SplitN(string(data), "\n", -1) // la fonction strings.SplitN permet de séparer le contenu du tableau à chaque saut de ligne. On crée ici un tableau de slice, chaque slice contient un mot.
	ToFind := tbWord[rand.Intn(len(tbWord))]         // la fonction rand.Intn permet de générer un nbr random, grâce à la fonction rand.Seed le nbre random change à chaque tour.
	Attempt := 10
	word := []rune(ToFind)
	nbrL := len(ToFind)/2 + 1
	var l string

	for i := 0; i < len(ToFind); i++ {
		if nbrL >= 0 {
			j := rand.Intn(len(ToFind))
			word[j] = '_'
			nbrL--
		}
	}

	fmt.Println(string(word))

	for Attempt >= 0 {
		fmt.Printf("you have %v attemps\n", Attempt)
		fmt.Printf("Propose a lower letter\n")
		fmt.Scan(&l)
		Error(l)

		for i := 0; i < len(ToFind); i++ {
			if l[0] == ToFind[i] {
				word[i] = rune(ToFind[i])
				fmt.Println(string(word))
				if string(word) == ToFind {
					fmt.Printf("You win\n")
					return
				}
			}
		}
	}
	fmt.Println("You lost, the word to find was %v \n", ToFind)
}
