package main

import "fmt"

func main() {
	var totalDiscount float64 = 0.0
	for {
		var price, discount float64
		fmt.Print("Enter product price (-1 for exit): ")
		fmt.Scan(&price)
		fmt.Print("Enter discount (-1 for exit): ")
		fmt.Scan(&discount)
		if price == -1 || discount == -1 {
			break
		} else if price > 0 && discount > 0 {
			if discount <= 0.3*price || discount < 2000 {
				totalDiscount += discount
			}
		} else {
			continue
		}
	}
	fmt.Println("Total discount: ", totalDiscount)
}
