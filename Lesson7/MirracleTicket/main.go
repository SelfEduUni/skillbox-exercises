package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Miracle tickets.")
	minTicketNumber := 100000
	maxTicketNumber := 999999
	counter := 0
	for i := minTicketNumber; i <= maxTicketNumber; i++ {
		if isMiracleTicket(i) {
			fmt.Println(i)
			counter++
		}
	}
	fmt.Println("Total miracle tickets:", counter)
}

func isMiracleTicket(i int) bool {
	left := i / 1000
	right := i % 1000
	miracle := 0
	temp := left
	for j := 2; j >= 0; j-- {
		miracle += (temp / int(math.Pow(10, float64(j)))) * int(math.Pow(10, float64(2-j)))
		temp %= int(math.Pow(10, float64(j)))
	}
	return miracle == right
}
