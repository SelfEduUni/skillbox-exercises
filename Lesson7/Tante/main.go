package main

import "fmt"

func main() {
	fmt.Println("Herringbone.")
	fmt.Print("Enter height: ")
	var height int
	fmt.Scan(&height)
	field := 2*height - 1
	for i := 0; i < height; i++ {
		star := (2*i + 1)
		left := (field - star) / 2
		right := left
		for j := 0; j < left; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < star; j++ {
			fmt.Print("*")
		}
		for j := 0; j < right; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
