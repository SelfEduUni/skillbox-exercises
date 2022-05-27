package main

import "fmt"

func main() {
	var val1, val2 int
	fmt.Print("Enter first value: ")
	fmt.Scan(&val1)
	fmt.Print("Enter second value: ")
	fmt.Scan(&val2)
	if val1 > val2 {
		fmt.Println("The minimum value is: ", val2)
	} else if val1 < val2 {
		fmt.Println("The minimum value is: ", val1)
	} else {
		fmt.Println("Values are equal")
	}
}
