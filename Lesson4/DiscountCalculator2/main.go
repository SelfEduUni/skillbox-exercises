package main

import "fmt"

func main() {
	var dayOfWeek int
	var invoiceAmount float64
	var total float64 = 0.0
	fmt.Print("Введите номер дня недели: ")
	fmt.Scan(&dayOfWeek)
	if dayOfWeek < 1 {
		fmt.Println("Неверный номер дня недели")
	} else if dayOfWeek > 7 {
		fmt.Println("Неверный номер дня недели")
	} else {
		fmt.Print("Введите сумму счета: ")
		fmt.Scan(&invoiceAmount)
		if dayOfWeek >= 6 {
			if invoiceAmount > 10000 {
				discount := invoiceAmount * 0.1
				total = invoiceAmount - discount
			} else {
				total = invoiceAmount
				fmt.Println("Скидка не предусмотрена")
			}
		}
		fmt.Println("Сумма счета: ", total)
	}
}
