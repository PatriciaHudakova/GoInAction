//create goroutines and explore scheduler behaviour
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //allocate max n.o. of logical processors for scheduler to use

	var wg sync.WaitGroup //wg is used to wait for program to finish
	wg.Add(2)             //add a count of two, one for each goroutine

	fmt.Println("Start Goroutines")

	go func() { //anonymous function, create goroutine
		defer wg.Done() // schedule call to done to update main

		for count := 0; count < 3; count++ { //display alphabet three times
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() { //anonymous function, create goroutine
		defer wg.Done() // schedule call to done to update main

		for count := 0; count < 3; count++ { //display alphabet three times
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait() //wait for goroutines to finish

	fmt.Println("\nTerminating Program")
}
