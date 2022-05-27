package main

import "fmt"

func main() {
	var dayOfWeek int
	fmt.Print("Введите номер дня недели: ")
	fmt.Scan(&dayOfWeek)
	var common string = "кофе, чай, яица, молоко, тосты, компот"
	if dayOfWeek < 1 {
		fmt.Println("Неверный номер дня недели")
	} else if dayOfWeek > 7 {
		fmt.Println("Неверный номер дня недели")
	} else if dayOfWeek == 1 {
		fmt.Println(common, "Овсяная каша, картошка с грибами, борщ")
	} else if dayOfWeek == 2 {
		fmt.Println(common, "Суп, солянка, куриный суп")
	} else if dayOfWeek == 3 {
		fmt.Println(common, "Манная каша, бифстроганоф, луковый суп")
	} else if dayOfWeek == 4 {
		fmt.Println(common, "Гречневая каша, утка по-пекински, салат")
	} else if dayOfWeek == 5 {
		fmt.Println(common, "Перловка, сосиски, Щи")
	} else if dayOfWeek == 6 {
		fmt.Println(common, "Пшенная каша, котлеты куриные, венегрет")
	} else if dayOfWeek == 7 {
		fmt.Println(common, "Мюсли, Котлеты по-киевски, оливье")
	} else {
		fmt.Println("Wrong number")
	}
}
