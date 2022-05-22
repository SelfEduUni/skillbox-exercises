package main

import "fmt"

func main() {
	var product1 float64
	var product2 float64
	var product3 float64

	fmt.Println("Enter the price of product 1: ")
	fmt.Scan(&product1)
	fmt.Println("Enter the price of product 2: ")
	fmt.Scan(&product2)
	fmt.Println("Enter the price of product 3: ")
	fmt.Scan(&product3)

	total := product1 + product2 + product3
	if total > 10000 {
		discount := total * 0.1
		total -= discount
	}
	fmt.Println("Total price is: ", total)
}
