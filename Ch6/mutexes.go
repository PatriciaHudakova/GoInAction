package main

import "fmt"
import "runtime"
import "sync"

var (
	Counter int            //var inc by all goroutines
	WG      sync.WaitGroup //wait for program termination
	mutex   sync.Mutex     //define critical section of code
)

func main() {
	WG.Add(2)

	go incCounter3(1)
	go incCounter3(2)

	WG.wait()
	fmt.Printf("Final counter: %d\\n", Counter)
}

func incCounter3(id int) {
	defer WG.Done()

	for i := 0; i < 2; i++ {
		mutex.Lock() // only allow one goroutine through this
		{
			value := Counter  //capture value of counter
			runtime.Gosched() //yield thread and be placed back into queue
			value++
			Counter = value //store value back into counter
		}
		mutex.Unlock()
	}
}
