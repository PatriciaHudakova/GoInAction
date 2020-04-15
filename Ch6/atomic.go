package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter2 int64
	wg3      sync.WaitGroup
)

func main() {
	wg3.Add(2)
	go incCounter2(1)
	go incCounter2(2)

	wg3.Wait()
	fmt.Println("Final counter: ", counter2)
}

func incCounter2(id int) {
	defer wg2.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter2, 1) //enforces one goroutine at a time
		runtime.Gosched()
	}
}
