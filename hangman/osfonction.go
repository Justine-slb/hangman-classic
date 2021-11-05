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

func WordHoleF(toFind string, nLshow int) []rune {
	tbWord := []rune(toFind)
	for n := 0; n < len(tbWord)-1; n++ {
		tbWord[n] = '_'
	}
	tbIndex := RandomIndex(nLshow, tbWord) // creation of a random index Array with the HideLetters function
	for i := 0; i < len(tbIndex); i++ {    // loop for browse the tbIndex index by index
		for j := 0; j < len(toFind); j++ { // loop to browse the toFind string
			if j == tbIndex[i] { // when the value of tbIndex = value of j
				tbWord[j] = rune(toFind[j]) // the function transforms "_" in the tbWord[j] by the value oh toFind[j]
			}
		}
	}
	return tbWord
}

func RandomIndex(nLToShow int, tbWord []rune) []int { // Array with random index all different
	rand.Seed(time.Now().UnixNano())
	var tbIndex []int

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
