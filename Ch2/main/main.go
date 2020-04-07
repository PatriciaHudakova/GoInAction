package main //executable

import (
	"log"
	"os"
	"github.com/PatriciaHudakova/GoInAction/Ch2/search"
	_ "github.com/PatriciaHudakova/GoInAction/Ch2/matchers" //no identifiers from matchers package are used
)

func init() { //all init functions to be called before the main function
	log.SetOutput(os.Stdout) //logger writes onto stderr device by default
}

func main() { //again, the entry point of the program is here
	search.Run("president") //perform search
}