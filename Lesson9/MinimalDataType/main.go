package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 int16
	fmt.Print("num1 = ")
	fmt.Scan(&num1)
	fmt.Print("num2 = ")
	fmt.Scan(&num2)

	result := int32(num1) * int32(num2)

	switch {
	case result >= 0 && result <= int32(math.MaxUint8):
		fmt.Printf("%d * %d = %d - you could save it in UInt8 format\n", num1, num2, result)
	case result < 0 && result >= int32(math.MinInt8):
		fmt.Printf("%d * %d = %d - you could save it in Int8 format\n", num1, num2, result)
	case result >= 0 && result <= int32(math.MaxUint16):
		fmt.Printf("%d * %d = %d - you could save it in UInt16 format\n", num1, num2, result)
	case result < 0 && result >= int32(math.MinInt16):
		fmt.Printf("%d * %d = %d - you could save it in Int16 format\n", num1, num2, result)
	case result >= 0:
		fmt.Printf("%d * %d = %d - you could save it in UInt32 format\n", num1, num2, result)
	case result < 0:
		fmt.Printf("%d * %d = %d - you could save it in Int32 format\n", num1, num2, result)
	default:
		fmt.Println("This format is not recognized.")
	}
}
