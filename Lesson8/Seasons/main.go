package main

import "fmt"

func main() {
outer:
	for {
		fmt.Print("Enter a month: ")
		var month string
		fmt.Scanln(&month)
		switch month {
		case "January", "February", "December", "Декабрь", "Январь", "Февраль":
			fmt.Println("The month is in season: Winter")
		case "April", "May", "March", "Март", "Апрель", "Май":
			fmt.Println("The month is in season: Spring")
		case "July", "August", "June", "Июль", "Июнь", "Август":
			fmt.Println("The month is in season: Summer")
		case "October", "November", "September", "Сентябрь", "Октябрь", "Ноябрь":
			fmt.Println("The month is in season: Autumn")
		default:
			fmt.Println("Invalid month")
			break outer
		}
	}
}
