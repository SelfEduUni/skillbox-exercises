package main

import (
	"fmt"
	"math"
)

func main() {
	var counter8 int = 0
	var counter16 int = 0

	for i := 0; i < int(math.MaxUint32); i++ {
		if uint8(i) == 0 {
			counter8++
		}
		if uint16(i) == 0 {
			counter16++
		}
	}
	fmt.Println("counter8:", counter8)
	fmt.Println("counter16:", counter16)
}
