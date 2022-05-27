package main

import "fmt"

func main() {
	var val1, val2 int
	fmt.Print("Enter first value: ")
	fmt.Scan(&val1)
	fmt.Print("Enter second value: ")
	fmt.Scan(&val2)
	if val1%val2 == 0 {
		fmt.Println("The first value is divisible by the second value.")
	} else if val2%val1 == 0 {
		fmt.Println("The second value is divisible by the first value.")
	} else {
		fmt.Println("The values are not divisible.")
	}
}
