package main

import "fmt"

func main() {
	// FIRST WAY TO DECLARE A MAP
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"black": "#000000",
		"white": "#ffffff",
	}

	// HOW TO ACCESS / ASSIGN VALUES INSIDE MAP
	colors["red"] = "red"
	fmt.Println(colors["red"])

	// HOW TO DELETE A VALUE OF A MAP
	delete(colors, "red")

	// HOW TO ITERATE OVER A MAP
	printMap(colors)

	// SECOND WAY TO DECLARE A MAP
	// var colors map[string]string

	// THIRD WAY TO DECLARE A MAP
	// colors := make(map[string]string)

	fmt.Println(colors)
}

// HOW TO ITERATE OVER A MAP
func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
