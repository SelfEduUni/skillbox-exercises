package main

import "fmt"

func main() {
	var breakWord string = "finish"
	for {
		fmt.Println("Egg")
		var word string
		fmt.Scan(&word)
		if word == breakWord {
			break
		}
	}
}
