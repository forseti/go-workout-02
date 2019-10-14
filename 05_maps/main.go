package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[int]string)
	// colors[10] = "#ffffff"

	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// delete(colors, 10)

	// fmt.Println(colors)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}

}
