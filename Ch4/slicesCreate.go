package main

import "fmt"

func main() {

	slice := make([]int, 3, 5)   //a slice of length 3, capacity 5
	slice = []int{1, 2, 3, 4, 5} //initialised with elements

	slice2 := []string{99: ""} //100 elements long, capacity and length both 100

	//0 length, 0 capacity, nil pointer
	var slice3 []int //nil slice - slice with no initialisation

	sliceMake := make([]int, 0) //empty slice using make()
	sliceLit := []int{}         //empty slice using a literal

	newSlice := slice[1:3] // new slice of length 2, capacity 4
	//if outside capacity ==> runtime error
	newSlice[1] = 35 //change first index of newSlice to 35 ==> affects slice

}
