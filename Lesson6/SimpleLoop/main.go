package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	for i := 0; i < number; i++ {
		fmt.Println(i)
	}
}
