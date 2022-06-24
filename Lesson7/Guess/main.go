package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("I will guess some number between 0 and 100")
	fmt.Println("Try guess it")
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(101)
	counter := 0
	for {
		var guess int
		fmt.Print("Your guess: ")
		fmt.Scanf("%d", &guess)
		counter++
		if guess == number {
			fmt.Println("You guessed it!")
			break
		} else if guess < number {
			fmt.Println("Too low")
		} else {
			fmt.Println("Too high")
		}
	}
	fmt.Println("You guessed it in", counter, "tries")
}
