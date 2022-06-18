package main

import "fmt"

func main() {
	password := "ilovego!"
	var suggestion string
	for {
		fmt.Print("Enter a password: ")
		fmt.Scan(&suggestion)
		if suggestion == password {
			break
		}
	}
}
