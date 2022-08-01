package main

import "fmt"

func main() {
	// Create a map with initial values
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)

	// Declare variable of type map
	// var colors2 map[string]string
	colors2 := make(map[string]string)
	fmt.Println(colors2)
	colors2["white"] = "#ffffff"
	colors2["101"] = "1010101"
	fmt.Println(colors2)
	delete(colors2, "white")
	fmt.Println(colors2)

	colors["white"] = "#ffffff"
	// Iterate through map
	printMap(colors)

}

func printMap(m map[string]string) {
	fmt.Println("Printing map")
	for k, v := range m {
		fmt.Println("Hex code for", k, "is", v)
	}
}
