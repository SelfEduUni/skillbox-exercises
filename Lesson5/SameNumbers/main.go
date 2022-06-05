package main

import "fmt"

func main() {
	var number1, number2, number3 int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&number1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&number2)
	fmt.Print("Введите третье число: ")
	fmt.Scan(&number3)
	if number1 == number2 || number1 == number3 || number2 == number3 {
		fmt.Println("Среди чисел есть одинаковые")
	} else {
		fmt.Println("Среди чисел нет одинаковых")
	}
}
