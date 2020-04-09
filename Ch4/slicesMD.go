package main

import "fmt"

func main() {
	slice := [][]int{{10}, {10, 100}}
	slice[1] = append(slice[1], 20)

	fmt.Println(slice)

	bigSlice := make([]int, 1e6) //24 bytes, no need for pointers
	bigSlice = boo(bigSlice)
}

func boo(bigSlice []int) []int {
	return bigSlice
}
