package hangman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Save(game Game) {
	data := Game{game.WordRoot, game.WordHole, game.Attempt, game.Tried, ""}
	dataBytes, err := json.Marshal(data) //json.Marshal() convert Struct into Byte data, This method takes object as a param and returned Bytes code
	file, err := os.Create("save.txt")   //create a file save.txt
	if err != nil {
		log.Fatal("Error", err)
	}
	defer file.Close()
	fmt.Fprintf(file, string(dataBytes)) //take the Json Marshal into the file text
	return
}

func Load() Game { // this function read the args in put in the terminal
	data2 := Game{}
	//if os.Args[1] == "--startWith" { // use a flag
	//if len(os.Args) > 1 {
	//if os.Args[2] == "save.txt" { // the name of the file to read
	content, err := ioutil.ReadFile("save.txt")
	data2 = Game{}
	err = json.Unmarshal(content, &data2) // convert Json Marshal in structure
	if err != nil {                       // if err content data, there is an error, so that print an error message
		log.Fatal(err)
	}
	fmt.Println(data2)
	return data2
}
