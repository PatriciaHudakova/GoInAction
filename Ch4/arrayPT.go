package main

import "fmt"

func main() {
	array := [5]*int{1: new(int), 2: new(int)} //initialise with integer pointers
	*array[1] = 10                             //assign value
	*array[2] = 20                             //assign value

	fmt.Println(array)
}
