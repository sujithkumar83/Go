package main

import "fmt"

func main() {
	//Method 1 to define
	// var colors map[string] string
	//Mthod2
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "#000000",
	}
	//Method 3
	//colors["white"]= "#ffffff"

	//delete a ntry
	delete(colors, "red")

	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("Hx code for", color, "is ", hex)
	}
}
