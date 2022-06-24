package main

import "fmt"

func main() {
	fmt.Println("Chess Board.")
	var height, width int
	fmt.Print("Enter height: ")
	fmt.Scan(&height)
	fmt.Print("Enter width: ")
	fmt.Scan(&width)
	counter := 1
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if counter%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
			counter++
		}
		fmt.Println()
	}
}
