package main

import "fmt"

func main() {
	var array [1e6]int //array of 8 megabytes
	// foo(array)   parse in the array into a function
	foo(&array) //parse a pointer to the array
}

func foo(array *[1e6]int) {
	fmt.Println(array)
}
