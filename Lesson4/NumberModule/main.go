package main

import "fmt"

func main() {
	var number float64
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	if number < 0 {
		number = -number
	}
	fmt.Println("Модуль числа:", number)
}
