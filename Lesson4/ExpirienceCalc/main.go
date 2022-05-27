package main

import "fmt"

func main() {
	var exp int
	fmt.Print("Enter your experience points: ")
	fmt.Scan(&exp)
	if exp < 0 {
		fmt.Println("Invalid input")
	} else if exp < 1000 {
		fmt.Println("You are a newbie with level 1.")
	} else if exp < 2500 {
		fmt.Println("You are a Junior with level 2.")
	} else if exp < 5000 {
		fmt.Println("You are a Senior with level 3.")
	} else {
		fmt.Println("You are a Master with level 4.")
	}
}
