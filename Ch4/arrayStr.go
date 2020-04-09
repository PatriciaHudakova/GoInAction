package main

import "fmt"

func main() {
	var array1 [5]string
	array2 := [5]string{"red", "blue", "green", "yellow", "pink"}
	array1 = array2 //array2 into array1

	fmt.Println(array1)
}
