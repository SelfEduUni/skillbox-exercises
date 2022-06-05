package main

import "fmt"

func main() {
	var ticketNumber int
	fmt.Print("Введите номер билета: ")
	fmt.Scan(&ticketNumber)
	if ticketNumber == 3425 {
		fmt.Println("Поздравляю!!! Вы вытащили счастливый билет!")
	} else {
		var firstPart, secondPart, leftDigit, rightDigit int
		firstPart = ticketNumber / 100
		secondPart = ticketNumber % 100
		leftDigit = secondPart / 10
		rightDigit = secondPart % 10
		secondPart = rightDigit*10 + leftDigit
		if firstPart == secondPart {
			fmt.Println("Вы вытащили зеркальный билет!")
		} else {
			fmt.Println("Вы вытащили обыкновенный билет.")
		}
	}
}
