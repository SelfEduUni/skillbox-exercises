package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Enter a first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter a second number: ")
	fmt.Scan(&num2)
	for i := 0; i < num2; i++ {
		num1++
	}
	fmt.Println(num1)
}
