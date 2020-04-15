package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var waitG sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix()) //seed random number generator
}

func main() {
	tasks := make(chan string, taskLoad) // create buffered channel

	waitG.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ { //launch gorotines to handle work
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d, post")
	}

	close(tasks) //close channel so goroutines will quit

	waitG.Wait()
}

func worker(tasks chan string, worker int) {
	defer waitG.Done()

	for {
		//wait for work to be assigned
		task, ok := <-tasks
		if !ok { //channel empty and closed
			fmt.Printf("worker: %d : shutting down \n", worker)
			return
		}

		fmt.Printf("worker %d started %s\n", worker, task)

		//randomly wait to simulate work time
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("worker %d completed %s\n", worker, task)
	}
}
