package main

import "fmt"

func main() {
	var cordX, cordY int
	fmt.Print("Введите координату X точки: ")
	fmt.Scan(&cordX)
	fmt.Print("Введите координату Y точки: ")
	fmt.Scan(&cordY)
	if cordX == 0 && cordY == 0 {
		fmt.Println("Точка в начале координат")
	} else if cordX > 0 && cordY > 0 {
		fmt.Println("Точка в первой четверти")
	} else if cordX < 0 && cordY < 0 {
		fmt.Println("Точка в третьей четверти")
	} else if cordX < 0 && cordY > 0 {
		fmt.Println("Точка во второй четверти")
	} else if cordX > 0 && cordY < 0 {
		fmt.Println("Точка в четвертой четверти")
	} else if cordX == 0 && cordY > 0 {
		fmt.Println("Точка на оси ОY в первой и второй четвертях")
	} else if cordX == 0 && cordY < 0 {
		fmt.Println("Точка на оси ОY в третьей и четвертой четверти")
	} else if cordX > 0 && cordY == 0 {
		fmt.Println("Точка на оси ОX в первой и четвертой четвертях")
	} else if cordX < 0 && cordY == 0 {
		fmt.Println("Точка на оси ОX во второй и третьей четверти")
	} else {
		fmt.Println("Не известно где находится точка")
	}
}
