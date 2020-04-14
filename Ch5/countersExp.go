package main

import "fmt"
import "github.com/PatriciaHudakova/GoInAction/Ch5/unexported"

func main() {
	counter := unexported.New(10)
	fmt.Println(counter)
}
