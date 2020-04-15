package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64 //flag to alert goroutines to shutdown
	wait     sync.WaitGroup
)

func main() {
	wait.Add(2) //add a count of two

	go doWork("A") //create two goroutines
	go doWork("B")

	time.Sleep(1 * time.Second) //give goroutines time to run

	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1) //safely flag it is time to shut down

	wait.Wait() //wait for goroutines to finish
}

func doWork(name string) { //checking shutdown flag to terminate early
	defer wait.Done()

	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s down\n", name)
			break
		}
	}
}
