package main

import (
	"fmt"
	"sync"
	"time"
)

var Wg sync.WaitGroup

func main() {
	baton := make(chan int) //create unbuffered channel

	Wg.Add(1) //count of one for the last runner

	go Runner(baton) //first runner to his mark

	baton <- 1 //race start

	Wg.Wait() //wait for race to finish
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton //wait to receive baton

	fmt.Printf("Runner %d running with baton\n", runner) //start running

	if runner != 4 { //new runner to the line
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond) //running around the track

	if runner == 4 { //race over
		fmt.Printf("Runner %d finished, race over\n", runner)
		Wg.Done()
		return
	}

	fmt.Printf("runner %d exchange with runner %d\n", runner, newRunner) //exchange of baton
	baton <- newRunner
}
