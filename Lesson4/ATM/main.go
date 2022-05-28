package main

import "fmt"

func main() {
	fmt.Println("Банкомат.")
	var cashwithdrawal int
	fmt.Println("Введите сумму для снятия:")
	fmt.Scan(&cashwithdrawal)
	if cashwithdrawal > 100000 {
		fmt.Println("Сумма превышает допустимый лимит снятия наличности.")
	} else if cashwithdrawal%100 != 0 {
		fmt.Println("Сумма должна быть кратна 100.")
	} else {
		fmt.Println("Операция успешно выполнена.")
		fmt.Println("Вы сняли", cashwithdrawal, "рублей.")
	}
}
