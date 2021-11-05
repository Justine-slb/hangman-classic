package hangman

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

func ChooseWord() string {
	rand.Seed(time.Now().UnixNano())
	f, e := os.Open("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\dictionnaire.txt")
	check(e)
	defer f.Close()
	data, _ := os.ReadFile("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\dictionnaire.txt")
	tbWords := strings.SplitN(string(data), "\n", -1)
	toFind := tbWords[rand.Intn(len(tbWords))]
	return toFind
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ArrayHangman() []string {
	f, e := os.Open("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\hangman.txt")
	check(e) // check erreur : par de fichier
	defer f.Close()
	data, err := os.ReadFile("C:\\Users\\justi\\OneDrive\\Documents\\YNOV\\YTRACK\\hangman-classic\\hangman.txt")
	check(err)
	tbHangman := strings.SplitN(string(data), "*", -1)
	return tbHangman
}

func PrintWord(toFind string, nLshow int) []rune {
	tbWord := []rune(toFind)
	for n := 0; n < len(tbWord)-1; n++ {
		tbWord[n] = '_'
	}
	tbIndex := RandomIndex(nLshow, tbWord) // création d'un tableau d'index aléatoire gràce à la fonction HideLetters
	for i := 0; i < len(tbIndex); i++ {    // on parcourt le tb index avec
		for j := 0; j < len(toFind); j++ { // on parcourt le tb toFind
			if j == tbIndex[i] { // quand la valeur comprise dans tbIndex[i] = index de toFind j'applique :
				tbWord[tbIndex[i]] = rune(toFind[tbIndex[i]]) //
			}

		}
	}
	return tbWord
}

func RandomIndex(nLToShow int, tbWord []rune) []int { // Cette fonction permet de créer un tableau d'index aléatoire, et tous différents
	rand.Seed(time.Now().UnixNano())
	var tbIndex []int
	for n := 0; n < len(tbWord)-1; n++ { // on commence par modifier tous les caractères du tbWord par des _.
		tbWord[n] = '_'
	}
	for i := 0; i < nLToShow; i++ {
		Rnbr := rand.Intn(len(tbWord) - 1)
		for j := 0; j < len(tbIndex); j++ {
			for Rnbr == tbIndex[j] { // dans cette boucle, si l'index random est déjà dans le tableau, on rappelle la focntion RandomIndex : Récursive
				return RandomIndex(nLToShow, tbWord)
			}
		}
		tbIndex = append(tbIndex, Rnbr) // si l'index random est valide : on l'ajoute au tbIndex.
	}
	return tbIndex
}
