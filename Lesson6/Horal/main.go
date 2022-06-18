package main

import "fmt"

func main() {
	var counter int = 0
	var incr int = 1
	for {
		if counter == 0 {
			fmt.Println("Check Point")
		} else {
			fmt.Println(counter)
		}
		if counter == 10 {
			incr = -1
		}
		if counter == -10 {
			incr = 1
		}
		counter += incr
	}
}
