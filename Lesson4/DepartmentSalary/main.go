package main

import "fmt"

func main() {
	var sal1, sal2, sal3 int
	fmt.Print("Введите зарплату первого сотрудника: ")
	fmt.Scan(&sal1)
	fmt.Print("Введите зарплату второго сотрудника: ")
	fmt.Scan(&sal2)
	fmt.Print("Введите зарплату третьего сотрудника: ")
	fmt.Scan(&sal3)

	sum := sal1 + sal2 + sal3
	average := float64(sum / 3)
	fmt.Printf("Средняя зарплата: %.2f\n", average)
	var max, min int
	max = sal1
	min = sal3
	if sal2 > max {
		max = sal2
	} else if sal2 < min {
		min = sal2
	}
	if sal3 > max {
		max = sal3
	} else if sal3 < min {
		min = sal3
	}
	if sal1 > max {
		max = sal1
	} else if sal1 < min {
		min = sal1
	}
	fmt.Println("Максимальная зарплата: ", max)
	fmt.Println("Минимальная зарплата: ", min)

}
