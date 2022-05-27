package main

import "fmt"

func main() {
	var cosm1, cosm2, cosm3 int
	fmt.Print("Введите уровень IQ первого космонавта: ")
	fmt.Scan(&cosm1)
	fmt.Print("Введите уровень IQ второго космонавта: ")
	fmt.Scan(&cosm2)
	fmt.Print("Введите уровень IQ третьего космонавта: ")
	fmt.Scan(&cosm3)

	if cosm1 >= cosm2 {
		if cosm1 >= cosm3 {
			fmt.Println("Первый космонавт назначается командиром")
		} else {
			fmt.Println("Третий космонавт назначается командиром")
		}
	} else {
		if cosm2 >= cosm3 {
			fmt.Println("Второй космонавт назначается командиром")
		} else {
			fmt.Println("Третий космонавт назначается командиром")
		}
	}
}
