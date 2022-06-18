package main

import "fmt"

func main() {
	var passengersInFloorFour, passengersInFloorSeven, passengersInFloorTen int = 3, 3, 3
	var elevatorCapacity int = 0
	var currentFloor int = 1
	var iteration int = 0
	var flag bool = false
	var index int = 1
	for {
		currentFloor += index
		iteration++
		if currentFloor == 24 {
			index = -1
			flag = true
			fmt.Println("Elevator is at floor 24")
			fmt.Println()
		}
		if currentFloor == 1 {
			index = 1
			flag = false
			fmt.Println(elevatorCapacity, "passengers left building at Floor 1 using elevator.")
			fmt.Println()
			elevatorCapacity = 0
			if passengersInFloorFour == 0 && passengersInFloorSeven == 0 && passengersInFloorTen == 0 {
				break
			}
		}
		if currentFloor == 10 && passengersInFloorTen != 0 && elevatorCapacity != 2 && flag {
			if elevatorCapacity == 0 {
				if passengersInFloorTen/2 != 0 {
					elevatorCapacity += 2
					passengersInFloorTen -= elevatorCapacity
				} else {
					elevatorCapacity += passengersInFloorTen
					passengersInFloorTen -= elevatorCapacity
				}
			} else {
				elevatorCapacity++
				passengersInFloorTen--
			}
			fmt.Println("Elevator is at floor 10 and is loading passengers...")
			fmt.Println(elevatorCapacity, "passengers are loaded into elevator.")
			fmt.Println(passengersInFloorTen, "passengers are left in floor 10.")
		}
		if currentFloor == 7 && passengersInFloorSeven != 0 && elevatorCapacity != 2 && flag {
			if elevatorCapacity == 0 {
				if passengersInFloorSeven/2 != 0 {
					elevatorCapacity += 2
					passengersInFloorSeven -= elevatorCapacity
				} else {
					elevatorCapacity += passengersInFloorSeven
					passengersInFloorSeven -= elevatorCapacity
				}
			} else {
				elevatorCapacity++
				passengersInFloorSeven--
			}
			fmt.Println("Elevator is at floor 7 and is loading passengers...")
			fmt.Println(elevatorCapacity, "passengers are loaded into elevator.")
			fmt.Println(passengersInFloorSeven, "passengers are left in floor 7.")

		}
		if currentFloor == 4 && passengersInFloorFour != 0 && elevatorCapacity != 2 && flag {
			if elevatorCapacity == 0 {
				if passengersInFloorFour/2 != 0 {
					elevatorCapacity += 2
					passengersInFloorFour -= elevatorCapacity
				} else {
					elevatorCapacity += passengersInFloorFour
					passengersInFloorFour -= elevatorCapacity
				}
			} else {
				elevatorCapacity++
				passengersInFloorFour--
			}
			fmt.Println("Elevator is at floor 4 and is loading passengers...")
			fmt.Println(elevatorCapacity, "passengers are loaded into elevator.")
			fmt.Println(passengersInFloorFour, "passengers are left in floor 4.")

		}
	}
	fmt.Println("Elevator made", iteration, "iterations.")
}
