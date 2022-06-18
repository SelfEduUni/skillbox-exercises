package main

import "fmt"

func main() {
	var sum int = 0
	for i := 1; i <= 40; i++ {
		if i%4 == 0 {
			sum += i
		} else {
			continue
		}
	}
	fmt.Println(sum)
}
