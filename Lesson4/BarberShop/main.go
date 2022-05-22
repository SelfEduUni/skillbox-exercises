package main

import "fmt"

func main() {
	// totalPeople := 18000
	// totalMen := 9000
	amountDaysInMonth := 30
	barberWorkTimePerDay := 8
	barberProductivity := 1
	// clientsPerDayStandard := totalMen / amountDaysInMonth
	var menInCity int
	var currentBarberInCity int
	fmt.Print("Сколько в городе мужчин? ")
	fmt.Scan(&menInCity)
	fmt.Print("Сколько в городе работает барбер-специалистов? ")
	fmt.Scan(&currentBarberInCity)
	clientsPerDay := menInCity / amountDaysInMonth
	clientsPerBarberInDay := barberProductivity * barberWorkTimePerDay
	if clientsPerDay <= clientsPerBarberInDay*currentBarberInCity {
		fmt.Println("Барбер-специалистов достаточно")
	}
	if clientsPerDay > clientsPerBarberInDay*currentBarberInCity {
		barberRequired := (clientsPerDay - clientsPerBarberInDay*currentBarberInCity) / (barberProductivity * barberWorkTimePerDay)
		fmt.Println("Необходимо найти более", barberRequired, "барбер-специалистов для продолжения работы в городе")
	}
}
