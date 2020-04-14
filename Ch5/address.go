package main

import "fmt"

type duration int

func (d *duration) pretty() string { //pointer receiver, returns a string
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	duration(42).pretty() //shouldn't be able to find the address
}
