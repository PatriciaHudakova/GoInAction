package main

import "fmt"

func main() {
	//var nil map[string]string //nil map, if values assigned ==> runtime error
	//dict := make(map[string]int)  // not idiomatic
	dict2 := map[string]string{"Red": "#da1337", "Orange": "e95a22"} //using a map literal
	//dict3 := map[int][]string{}  //map using slice of strings as value

	colours := map[string]string{}
	colours["Red"] = "#da1337"

	//testing by checking for value and flag
	value, exists := colours["Red"]
	if exists {
		fmt.Println(value)
	}

	//testing for a zero value if key does not exist
	value2 := dict2["Red"]
	if value2 != "" {
		fmt.Println(value2)
	}
}
