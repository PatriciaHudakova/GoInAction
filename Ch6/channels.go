package main

func main() {
	unbuffered := make(chan int) //unbuffered channel of ints

	buffered := make(chan string, 10) //buffered channel of strings

	buffered <- "Gopher" //send a string through the channel

	value := <-buffered //receive a string from the channel

}
