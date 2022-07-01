package main

import "fmt"

func main() {
	fmt.Print("Enter a day of the week in short format: ")
	var day string
	fmt.Scanln(&day)
	switch {
	case day == "Mon":
		fmt.Println("Monday")
		fallthrough
	case day == "Tue":
		fmt.Println("Tuesday")
		fallthrough
	case day == "Wed":
		fmt.Println("Wednesday")
		fallthrough
	case day == "Thu":
		fmt.Println("Thursday")
		fallthrough
	case day == "Fri":
		fmt.Println("Friday")
	default:
		fmt.Println("Not a valid day")
	}
}
