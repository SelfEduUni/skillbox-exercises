package main

import "fmt"

func main() {
	var sum float64 = 0.0
	for i := 0; i < 4; i++ {
		var number float64
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)
		sum += number
	}
	fmt.Println("The sum is: ", sum)
}
