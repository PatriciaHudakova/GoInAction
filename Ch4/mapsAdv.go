package main

import "fmt"

func main() {
	colours := map[string]string{
		"AliceBlue":   "f0f8ff",
		"Coral":       "ff7F50",
		"DarkGrey":    "a9a9a9",
		"ForestGreen": "228b22",
	}

	// display map using range
	for key, value := range colours {
		fmt.Println(key, value)
	}

	removeColours(colours, "Coral")

	for key, value := range colours {
		fmt.Println(key, value)
	}
}

func removeColours(colours map[string]string, key string) { //delete map values
	delete(colours, "Coral")
}
