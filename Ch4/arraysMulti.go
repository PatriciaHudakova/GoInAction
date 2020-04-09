package main

import "fmt"

func main() {
	var array [2][2]int                       //4x2 multi array
	array2 := [2][2]int{{1, 3}, {4, 7}}       //declare and initialise all elements
	array3 := [4][2]int{1: {1, 3}, 2: {4, 7}} //declare and initialise index 1 and 3
	array4 := [4][2]int{1: {0: 3}, 2: {1: 7}} //declare and initialise inner and outer array

	array = array2
	array3 = array4

	var value int = array2[0][1]

	fmt.Println(value)
	fmt.Println(array)
	fmt.Println(array3)
}
