package main

import "fmt"

func main() {
	dayOne := []int{5, 5, 5, 10, 20}
	dayTwo := []int{10, 10}
	dayThree := []int{5, 5, 10, 10, 20}
	dayFour := []int{5, 5, 10}
	dayFive := []int{5, 5, 10, 10, 20, 20, 20, 20, 20, 20}
	fmt.Println("Day 1: ", lemonadeChange(dayOne))
	fmt.Println("Day 2: ", lemonadeChange(dayTwo))
	fmt.Println("Day 3: ", lemonadeChange(dayThree))
	fmt.Println("Day 4: ", lemonadeChange(dayFour))
	fmt.Println("Day 5: ", lemonadeChange(dayFive))
}

func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five == 0 {
				return false
			}
			ten++
			five--
		case 20:
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
