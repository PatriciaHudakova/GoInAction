package search //performs search logic

import (
	"log"  //provides support for logging messages into stream
	"sync" //provides support for synchronising goroutines
)

//map of matchers for searching
var matchers = make(map[string]Matcher) //package-level variable, unexported

func Run(searchTerm string) { //performs search logic
	//get a list of feeds to search through
	feeds, err := RetrieveFeeds() // := declares and initialises
	if err != nil {
		log.Fatal(err)
	}

	//unbuffered channel to get match results
	results := make(chan *Result) //channels implement a queue

	//wait group
	var waitGroup sync.WaitGroup

	//set number of goroutines, they process individual feeds
	waitGroup.Add(len(feeds))

	//launch goroutine for each feed to find the results
	for _, feed := range feeds { // returns index and value if "_" not present
		//retrieve matcher for search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		//launch goroutine to run the search
		go func(matcher Matcher, feed *Feed) { //anonymous independent goroutine function, feed is a pointer variable
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	//goroutine to monitor when all work is done
	go func() {
		//wait for all to be processed
		waitGroup.Wait()

		//close the channel and exit program
		close(results)
	}()

	//display results as they come and return after the final result has been displayed
	Display(results)
}
