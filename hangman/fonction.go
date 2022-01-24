package hangman

import (
	"fmt"
	"strings"
)

// Game a game struct
type Game struct {
	WordRoot string
	WordHole []rune
	Attempt  int
	Tried    []string
	Input    string
}

func Level() int {
	var level int
	fmt.Println("Choose your level:")
	fmt.Scan(&level)
	return level
}

func Init(level int) (Game, int) {
	game := Game{}
	game.WordRoot = ChooseWord(level)
	game.WordHole = WordHoleF(game.WordRoot)
	game.Attempt = 10
	game.Tried = []string{}
	return game, level
}

// GameManager define params to play
func GameManager(game Game, level int) {
	win := GameHangman(game, level)
	if win {
		fmt.Println("You Win!!")
	} else {
		fmt.Println("Looser!!")
	}
	Menu()
}

// Counter update game counter
func Counter(attempt int, win bool, loose int) int {
	if win {
		return attempt
	}
	return attempt - loose
}

// Tested Check for previously tested chars
func Tested(game Game) ([]string, bool) {
	for _, char := range game.Tried { // compare with the letter already tried
		if game.Input == char {
			fmt.Println("you already tried this!")
			return game.Tried, false
		}
	}
	newLTried := append(game.Tried, game.Input)
	return newLTried, true
}

// PrintGame print game state on standard output
func PrintGame(game Game) {
	fmt.Println(PrintHangman(game.Attempt))
	fmt.Println(string(game.WordHole))
	fmt.Printf("You have %d attempt\n", game.Attempt)
}

func GameHangman(game Game, level int) bool {
	var found, validChar, tried bool
	for game.Attempt > 0 {
		PrintGame(game)
		game.Input = Input()
		if game.Input == "--SAVE" {
			Save(game, level)
			return found
		}
		if IsAlpha(game.Input) == false {
			GameHangman(game, level)
		}
		game.Tried, tried = Tested(game)
		if tried == false {
			GameHangman(game, level)
		}
		if len(game.Input) == 1 { // if it's a single character
			validChar, found = InputChar(game.Input, game.WordRoot, game.WordHole)
			game.Attempt = Counter(game.Attempt, validChar, 1)
		} else if len(game.Input) == len(game.WordRoot)-1 {
			found = InputWord(game.Input, game.WordRoot)
			game.Attempt = Counter(game.Attempt, found, 2)
		} else {
			fmt.Println("Your input is not conform, retry please")
			GameHangman(game, level)
		}
		//game.Attempt = Counter(game.Attempt, validChar)
		if found == true {
			return found
		}
	}
	return found
}

// Input get a new input and directly format it to upper case
func Input() string {
	var input string
	fmt.Scan(&input)
	input = strings.ToUpper(input)
	return input
}

func IsAlpha(s string) bool { // check if the input is only upper alpha character
	for _, value := range []rune(s) {
		if (value >= 'A' && value <= 'Z') || (value >= 'a' && value <= 'z') {
			return true // else return true
		}
	}
	fmt.Println("Only letter please")
	return false // return false if the input is not alpha
}

func InputChar(input, toFind string, wordHole []rune) (bool, bool) { // if the lengths of the input is 1 so a single letter
	validChar := false
	found := false
	for i := 0; i < len(toFind); i++ {
		if input == string(toFind[i]) { // if the input is in the word to found
			if wordHole[i] == '_' { // if the letter is hide yet
				wordHole[i] = rune(toFind[i])          // the letter is show
				validChar = true                       // the bool win letter take the true value
				newWordHole := string(wordHole)        // we update the wordHole
				found = InputWord(newWordHole, toFind) // we call the function InPutWord, to check is the word is finding. The function return a bol and the result is keeping int the bool Find
			}
		}
	}
	return validChar, found // the function return the 2 update bool
}

//Necessaire? fait le meme boulot que Compare()
// InputWord check if the current word equals the secret
func InputWord(input, toFind string) bool {
	different := Compare(input, toFind)
	var found bool
	if !different {
		fmt.Println(toFind)
		found = true // if there is no difference, input = toFind , find == true
	}
	return found // return update found
}

// Compare check index by index if the input and secret word are equals
func Compare(input, toFind string) bool {
	different := false
	for indexA := range input {
		for i := 0; i < len(toFind); i++ {
			if input[indexA] != toFind[indexA] {
				different = true // if there is a difference, different = true
				return different // return true in the function InputWord
			}
		}
	}
	return different // return update different to function InputWord
}

func PrintHangman(attempt int) string { // function to print the good hangman step
	HangmanState := ArrayHangman()
	state := 10 - attempt
	return HangmanState[state]
}
