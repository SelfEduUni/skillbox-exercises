package main

import "fmt"

func main() {
	var num1, num2, num3 int = 0, 0, 0
	for {
		if num1 < 10 {
			num1++
		}
		if num2 < 100 {
			num2++
		}
		if num3 == 1000 {
			break
		}
		num3++
	}
	fmt.Println("First number:", num1, "Second number", num2, "Third number", num3)
}
