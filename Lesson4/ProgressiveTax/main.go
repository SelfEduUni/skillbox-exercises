package main

import "fmt"

func main() {
	var income float64
	fmt.Println("Введите ваш доход:")
	fmt.Scan(&income)
	var tax float64 = 0.0
	if income > 50000 {
		tax = (income-50000)*0.3 + (50000-10000)*0.2 + (10000-0)*0.13
	} else if income > 10000 {
		tax = (income-10000)*0.2 + (10000-0)*0.13
	} else if income > 0 {
		tax = (income - 0) * 0.13
	} else {
		fmt.Println("Вы не верно указали свой доход.")
	}
	fmt.Println("Ваш налог к уплате:", tax)
	fmt.Println("Ваш доход после вычета налога:", income-tax)
}
