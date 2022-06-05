package main

import "fmt"

func main() {
	var rate1, rate2, rate3 float64
	fmt.Print("Введите первую ставку: ")
	fmt.Scan(&rate1)
	fmt.Print("Введите вторую ставку: ")
	fmt.Scan(&rate2)
	fmt.Print("Введите третью ставку: ")
	fmt.Scan(&rate3)
	if rate1 == rate2 && rate2 == rate3 {
		fmt.Println("Все ставки равны")
	} else {
		if rate1 <= rate2 && rate1 <= rate3 {
			if rate2 <= rate3 {
				fmt.Println("Ставки", rate1, "и", rate2, "наиболее выгодные.")
			} else {
				fmt.Println("Ставки", rate1, "и", rate3, "наиболее выгодные.")
			}
		} else {
			if rate2 <= rate1 && rate2 <= rate3 {
				if rate1 <= rate3 {
					fmt.Println("Ставки", rate2, "и", rate1, "наиболее выгодные.")
				} else {
					fmt.Println("Ставки", rate2, "и", rate3, "наиболее выгодные.")
				}
			} else {
				if rate3 <= rate1 && rate3 <= rate2 {
					if rate1 <= rate2 {
						fmt.Println("Ставки", rate3, "и", rate1, "наиболее выгодные.")
					} else {
						fmt.Println("Ставки", rate3, "и", rate2, "наиболее выгодные.")
					}
				}
			}
		}
	}
}
