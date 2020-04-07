package search

import (
	"fmt"
	"log"
)

type Result struct {
	Field   string
	Content string
}

type Matcher interface { //defines behaviour required by types that want to implement a new search type
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//launched as goroutine for each feed to ensure concurrency
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm) //perform search against a matcher using values and pointers
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) { //write to terminal as results are received
	for result := range results { //loop terminates when channel is closed
		fmt.Printf("%s: \n%s\n\n", result.Field, result.Content)
	}
}
