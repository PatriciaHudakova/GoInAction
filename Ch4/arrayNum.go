package main

import "fmt"

//var array [5]int  all values initialised to 0
// array := [...]int{1,2,3,4,5} //go will identify length after initialisation
// array := [5]int{1: 2, 2: 3}  initialise specific values

func main() {
	array := [5]int{1, 2, 3, 4, 5} //declare and initialise
	array[2] = 4
	fmt.Println(array)
}
