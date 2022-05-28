package main

import "fmt"

func main() {
	var dayOfWeek, numberOfGuests int
	var invoiceAmount float64
	fmt.Println("Введите день недели:")
	fmt.Scan(&dayOfWeek)
	fmt.Println("Введите количество гостей:")
	fmt.Scan(&numberOfGuests)
	fmt.Println("Введите сумму чека:")
	fmt.Scan(&invoiceAmount)
	var total float64 = 0.0
	var discount, extraCharge float64 = 0.0, 0.0
	if dayOfWeek < 1 {
		fmt.Println("Неверный номер дня недели")
	} else if dayOfWeek > 7 {
		fmt.Println("Неверный номер дня недели")
	} else if dayOfWeek == 1 {
		discount = 0.1
		fmt.Println("Скидка по понедельникам:", invoiceAmount*discount)
	} else if dayOfWeek == 5 {
		if invoiceAmount > 10000 {
			discount = 0.05
			fmt.Println("Скидка по пятницам:", invoiceAmount*discount)
		}
	}
	if numberOfGuests < 1 {
		fmt.Println("Неверное количество гостей")
	} else if numberOfGuests > 5 {
		extraCharge = 0.1
		fmt.Println("Надбавка на обслуживание:", invoiceAmount*extraCharge)
	}
	total = invoiceAmount * (1 - discount + extraCharge)
	fmt.Println("Сумма к оплате:", total)
}
