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
	dataBytes, err := json.Marshal(data) //json.Marshal() convert Struct into Byte data, This method takes object as a param and returned Bytes code
	file, err := os.Create("save.txt")   //create a file save.txt
	if err != nil {
		log.Fatal("Error", err)
	}
	defer file.Close()
	fmt.Fprintf(file, string(dataBytes)) //take the Json Marshal into the file text
	return
}

func Load() { // this function read the args in put in the terminal
	if os.Args[1] == "--startWith" { // use a flag
		if len(os.Args) > 1 {
			if os.Args[2] == "save.txt" { // the name of the file to read
				content, err := ioutil.ReadFile("save.txt")
				data2 := game{}
				err = json.Unmarshal(content, &data2) // convert Json Marshal in structure
				if err != nil {                       // if err content data, there is an error, so that print an error message
					log.Fatal(err)
				}
				GameHangman(data2.Word, data2.ToFind, data2.Attempt, data2.LTried)
			}
		}
	}
	return
}
