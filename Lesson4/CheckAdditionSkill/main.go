package main

import "fmt"

func main() {
	var val1, val2 int
	fmt.Print("Enter first value: ")
	fmt.Scan(&val1)
	fmt.Print("Enter second value: ")
	fmt.Scan(&val2)
	result := val1 + val2
	var guess int
	fmt.Print("Guess the result: ")
	fmt.Scan(&guess)
	if guess == result {
		fmt.Println("You guessed it!")
	} else {
		fmt.Println("You guessed wrong! The result is:", result)
	}
}
