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
	var flag bool = false
	if number1 > 5 {
		flag = true
	}
	if number2 > 5 {
		flag = true
	}
	if number3 > 5 {
		flag = true
	}
	if flag {
		fmt.Println("Среди введенных чисел есть число больше 5.")
	} else {
		fmt.Println("Среди введенных чисел нет числа больше 5.")
	}
}
