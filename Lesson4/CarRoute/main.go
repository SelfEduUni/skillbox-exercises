package main

import "fmt"

func main() {
	var speed int
	fmt.Print("Введите скорость автомобиля: ")
	fmt.Scan(&speed)
	var distance int = 200
	time := float64(distance) / float64(speed)
	if time <= 2 {
		fmt.Println("Вы приехали.")
	} else {
		fmt.Println("Вы не доехали.")
	}
}
