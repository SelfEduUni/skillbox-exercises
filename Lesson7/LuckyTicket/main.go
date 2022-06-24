package main

import "fmt"

func main() {
	minTicketNumber := 100000
	maxTicketNumber := 999999
	previousTicketNumber := 0
	minDistance := 0
	maxDistance := 0
	distance := 0
	for i := minTicketNumber; i <= maxTicketNumber; i++ {
		left := i / 1000
		right := i % 1000
		if digitSum(left) == digitSum(right) {
			if previousTicketNumber == 0 {
				previousTicketNumber = i
				continue
			}
			distance = i - previousTicketNumber
			if distance < minDistance {
				minDistance = distance
			}
			if distance > maxDistance {
				maxDistance = distance
				previousTicketNumber = i
			}
		}
	}
	fmt.Println("To be absolutely sure that you have lucky ticket,")
	fmt.Println("you need buy ticket", maxDistance, "tickets.")
}

func digitSum(num int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += num % 10
		num /= 10
	}
	return sum
}
