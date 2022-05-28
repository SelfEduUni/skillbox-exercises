package main

import "fmt"

func main() {
	var exam1, exam2, exam3 int
	fmt.Println("Введите результат первого экзамена:")
	fmt.Scan(&exam1)
	fmt.Println("Введите результат второго экзамена:")
	fmt.Scan(&exam2)
	fmt.Println("Введите результат третьего экзамена:")
	fmt.Scan(&exam3)
	result := exam1 + exam2 + exam3
	fmt.Println("Сумма проходных баллов:", 275)
	fmt.Println("Количество набранных баллов:", result)
	if result >= 275 {
		fmt.Println("Вы поступили!")
	} else {
		fmt.Println("Вы не поступили.")
	}
}
