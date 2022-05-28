package main

import "fmt"

func main() {
	var number1, number2, number3 int
	fmt.Println("Три числа.")
	fmt.Println("Введите первое число:")
	fmt.Scan(&number1)
	fmt.Println("Введите второе число:")
	fmt.Scan(&number2)
	fmt.Println("Введите третье число:")
	fmt.Scan(&number3)
	var counter int = 0
	if number1 >= 5 {
		counter++
	}
	if number2 >= 5 {
		counter++
	}
	if number3 >= 5 {
		counter++
	}
	if counter > 0 {
		fmt.Println("Среди введенных чисел есть", counter, "число(а) больше или равны 5.")
	} else {
		fmt.Println("Среди введенных чисел нет чисел больше или равных 5.")
	}
}
