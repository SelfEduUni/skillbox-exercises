package main

import "fmt"

func main() {
	amountDaysInMonth := 30
	barberWorkTimePerDay := 8
	barberProductivity := 1
	var menInCity int
	var currentBarberInCity int
	fmt.Print("Сколько в городе мужчин? ")
	fmt.Scan(&menInCity)
	fmt.Print("Сколько в городе работает барбер-специалистов? ")
	fmt.Scan(&currentBarberInCity)
	clientsPerDay := menInCity / amountDaysInMonth
	if menInCity%amountDaysInMonth != 0 {
		clientsPerDay++
	}
	clientsPerBarberInDay := barberProductivity * barberWorkTimePerDay
	if clientsPerDay <= clientsPerBarberInDay*currentBarberInCity {
		fmt.Println("Барбер-специалистов достаточно")
	}
	if clientsPerDay > clientsPerBarberInDay*currentBarberInCity {
		barberRequired := (clientsPerDay - clientsPerBarberInDay*currentBarberInCity) / (barberProductivity * barberWorkTimePerDay)
		if (clientsPerDay-clientsPerBarberInDay*currentBarberInCity)%(barberProductivity*barberWorkTimePerDay) != 0 {
			barberRequired++
		}
		fmt.Println("Необходимо найти более", barberRequired, "барбер-специалистов для продолжения работы в городе")
	}
}
