package main

import (
	"fmt"
	"github.com/goinaction/code/chapter3/words"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1] //assign and declare variable filename

	contents, err := ioutil.ReadFile(filename) //if there is an error, print it to the terminal
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
