package main

import "fmt"
import "github.com/PatriciaHudakova/GoInAction/Ch5/unexported"

func main() {
	a := unexported.Admin{
		Rights: 10,
	}

	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Println(a)
}
