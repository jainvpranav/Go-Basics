package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
		"grey":  "#cccccc",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + "\t" + hex)
	}
}
