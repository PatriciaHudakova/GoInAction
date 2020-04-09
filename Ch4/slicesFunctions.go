package main

import "fmt"

func main() {
	slice := make([]string, 4, 5)             //make slice length 4, capacity 5
	slice = []string{"a", "b", "c", "d", "e"} //initialise slice elements
	newSlice := slice[2:5]                    //new slice c, d, e
	newSlice = append(newSlice, "f")          //overwrites length+1's element and increases length
	newSlice = append(slice, "g")             //doubles array size if <1000, otherwise +25%, appends next element to "g"

	fmt.Println(append(slice, newSlice...)) //variadic function

	for index, value := range slice { //iterating over slices
		fmt.Println(index, value)
	}

	for _, value := range slice { //black identifier to omit index
		fmt.Println(value)
	}

	for index := 2; index < len(slice); index++ { //traditional for loop for iteration
		fmt.Println(index, slice[index])
	}
}
