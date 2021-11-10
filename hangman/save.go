package hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type game struct {
	Word    []rune
	ToFind  string
	Attempt int
	LTried  []string
}

func Save(word []rune, toFind string, attempt int, lTried []string) {

	data := game{word, toFind, attempt, lTried}
	dataBytes, err := json.Marshal(data) //permet de convertir la structure en format json marshal
	file, err := os.Create("save.txt")   //cree un fichier save.txt
	if err != nil {
		log.Fatal("Error", err)
	}
	defer file.Close()
	fmt.Fprintf(file, string(dataBytes)) //place le json marshal dans le fichier text
	return
}

func Load() {
	//condition a placer dans le fichier main
	if os.Args[1] == "--startWith" {
		if len(os.Args) > 1 {
			if os.Args[2] == "save.txt" {
				content, err := ioutil.ReadFile("save.txt")
				data2 := game{}
				err = json.Unmarshal(content, &data2) //reconvertir le json marshall en structure
				if err != nil {
					log.Fatal(err)
				}
				GameHangman(data2.Word, data2.ToFind, data2.Attempt, data2.LTried)
			}
		}
	}
	return
}
