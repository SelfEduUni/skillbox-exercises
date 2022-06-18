package main

import "fmt"

func main() {
	var number, power int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	fmt.Print("Enter a power: ")
	fmt.Scan(&power)
	var result int = 1
	if power < 0 {
		fmt.Println("We do not yet power negative numbers.")
	}
	if power > 0 {
		for i := 0; i < power; i++ {
			result *= number
		}
	}
	fmt.Println("The result is: ", result)
}
