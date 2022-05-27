package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number (between 1 and 5): ")
	fmt.Scan(&number)
	guess := number / 2
	if guess == 0 {
		guess = 5
	}
	fmt.Println("I guess your number is", guess) // 1 or 5
	if number == guess {
		fmt.Println("I guessed correctly!")
	} else {
		var more int
		fmt.Print("Is it bigger or less (1/-1): ")
		fmt.Scan(&more)
		if more == 1 {
			guess++
			fmt.Println("I guess your number is", guess) // 2
			if guess == number {
				fmt.Println("I guessed correctly!")
			} else {
				fmt.Print("Is it bigger (1): ")
				fmt.Scan(&more)
				if more == 1 {
					guess++
					fmt.Println("I guess your number is", guess) // 4
					if guess == number {
						fmt.Println("I guessed correctly!")
					} else {
						guess++
						fmt.Println("I guess your number is", guess, "since this only last option.") // 5
					}
				}
			}
		} else {
			guess--
			fmt.Println("I guess your number is", guess) // 4
			if number == guess {
				fmt.Println("I guessed correctly!")
			} else {
				fmt.Print("Is it less (1): ")
				fmt.Scan(&more)
				if more == 1 {
					guess--
					fmt.Println("I guess your number is", guess) // 3
					if number == guess {
						fmt.Println("I guessed correctly!")
					} else {
						fmt.Print("Is it less (1): ")
						fmt.Scan(&more)
						if more == 1 {
							guess--
							fmt.Println("I guess your number is", guess) // 2
							if number == guess {
								fmt.Println("I guessed correctly!")
							} else {
								guess--
								fmt.Println("I guess your number is", guess, "since this only last option.") // 1
							}
						}
					}
				}
			}
		}
	}
}
