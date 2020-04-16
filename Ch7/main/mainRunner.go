package main

import (
	"github.com/PatriciaHudakova/GoInAction/Ch7"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second //3 secs

func main() {
	log.Println("Starting work")

	r := runner.New(timeout) //create a new runner

	r.Add(createTask(), createTask(), createTask()) //add runnable tasks

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}
	log.Println("Process Ended")
}

func createTask() func(int) { //returns example task sleeping for specified time
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
