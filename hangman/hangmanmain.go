package hangman

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

func HangmanMain() {
	rand.Seed(time.Now().UnixNano()) // fonction qui permet de générer un nouveau nombre random à chaque tour.
	f, e := os.Open("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\dictionnaire.txt")
	check(e)
	defer f.Close()
	data, _ := os.ReadFile("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\dictionnaire.txt")

	tbWords := strings.SplitN(string(data), "\n", -1) // la fonction strings.SplitN permet de séparer le contenu du tableau à chaque saut de ligne. On crée ici un tableau de slice, chaque slice contient un mot.

	toFind := tbWords[rand.Intn(len(tbWords))] // la fonction rand.Intn permet de générer un nbr random, grâce à la fonction rand.Seed le nbre random change à chaque tour.

	Attempt := 10

	nbrLShow := len(toFind)/2 - 1

	var sugestL string

	tbWord := []rune(toFind)

	var retry bool

	var tbSugestTry []string
	for n := 0; n < len(tbWord); n++ {
		tbWord[n] = '_'
	}

	fmt.Println(string(tbWord))
	tbWord = CacheLettre(toFind, tbWord, nbrLShow)
	compteur := 0

	for compteur < Attempt {
		fmt.Printf("you have %v attemps\n", Attempt)
		fmt.Println(string(tbWord))
		fmt.Printf("Propose a lower letter\n")
		fmt.Scan(&sugestL)
		Error(sugestL)
		tbSugestTry = append(tbSugestTry, sugestL)
		fmt.Println(tbSugestTry)

		for i := 0; i < len(toFind); i++ {
			if sugestL[0] != toFind[i] {
				retry = true
			} else {
				fmt.Println("Well done!")
				tbWord[i] = rune(toFind[i])
				fmt.Println(string(tbWord))
				if string(tbWord) != toFind {
					continue
				} else {
					fmt.Printf("You win\n")
					return
				}
			}
		}
		if retry == true {
			PrintHangman(compteur)
			compteur++
		}
	}
	fmt.Printf("You lost, the word to find was %v \n", toFind)
}

/*
func CacheLettre(toFind string, tbWord []rune, nLshow int) []rune {
	tb := HideLetters(nLshow, tbWord)
	for i := 0; i < len(tb); i++ {
		for j := 0; j < len(toFind); j++ {
			if j == tb[i] {
				tbWord[tb[i]] = rune(toFind[tb[i]])
			}
		}
	}
	return tbWord
}

func HideLetters(nLToShow int, tbWord []rune) []int {
	rand.Seed(time.Now().UnixNano())
	var tb []int

	for i := 0; i < nLToShow; i++ {
		Rnbr := rand.Intn(len(tbWord) - 1)
		for j := 0; j < len(tb); j++ {
			for Rnbr == tb[j] {
				return HideLetters(nLToShow, tbWord)
			}
		}
		tb = append(tb, Rnbr)
	}
	return tb
}

func PrintHangman ( a int ) { // fonction qui permet de générer un nouveau nombre random à chaque tour.
	f, e := os.Open("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\hangman.txt")
	check(e)
	defer f.Close()
	data, _ := os.ReadFile("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\hangman.txt")
	tbHangman := strings.SplitN(string(data), "*", -1)
	fmt.Println(tbHangman[a])
}

*/
