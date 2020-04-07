package search

import (
	"encoding/json" //support for encoding and decoding JSON
	"os"            //support for accessing OS functionality
)

const dataFile = "data/data.json"

type Feed struct { //declare struct type
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"site"`
}

func RetrieveFeeds() ([]*Feed, error) { // reads and decodes the data file, returns a pointer and an error
	file, err := os.Open(dataFile) //open file
	if err != nil {
		return nil, err
	}

	defer file.Close() //close file after function return, even after unexpected termination

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds) //decode files into slices of pointers

	return feeds, err
}
