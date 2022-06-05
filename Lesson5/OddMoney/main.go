package main

import "fmt"

func main() {
	fmt.Println("Банкомат приветствует Вас!")
	var cashValue1, cashValue2, cashValue3 int
	fmt.Print("Введите номинал купюры, хранимый в банкомате (1): ")
	fmt.Scan(&cashValue1)
	fmt.Print("Введите номинал купюры, хранимый в банкомате (2): ")
	fmt.Scan(&cashValue2)
	fmt.Print("Введите номинал купюры, хранимый в банкомате (3): ")
	fmt.Scan(&cashValue3)
	var cashWithdrawal int
	fmt.Print("Введите сумму для снятия с банкомата: ")
	fmt.Scan(&cashWithdrawal)
	var value1Counter, value2Counter, value3Counter, rest int
	if cashValue1 >= cashValue2 && cashValue1 >= cashValue3 {
		if cashValue2 >= cashValue3 {
			value1Counter = cashWithdrawal / cashValue1
			value2Counter = (cashWithdrawal - value1Counter*cashValue1) / cashValue2
			value3Counter = (cashWithdrawal - value1Counter*cashValue1 - value2Counter*cashValue2) / cashValue3
		} else {
			value1Counter = cashWithdrawal / cashValue1
			value3Counter = (cashWithdrawal - value1Counter*cashValue1) / cashValue3
			value2Counter = (cashWithdrawal - value1Counter*cashValue1 - value3Counter*cashValue3) / cashValue2
		}
	} else if cashValue2 >= cashValue3 && cashValue2 >= cashValue1 {
		if cashValue1 >= cashValue3 {
			value2Counter = cashWithdrawal / cashValue2
			value1Counter = (cashWithdrawal - value2Counter*cashValue2) / cashValue1
			value3Counter = (cashWithdrawal - value2Counter*cashValue2 - value1Counter*cashValue1) / cashValue3
		} else {
			value2Counter = cashWithdrawal / cashValue2
			value3Counter = (cashWithdrawal - value2Counter*cashValue2) / cashValue3
			value1Counter = (cashWithdrawal - value2Counter*cashValue2 - value3Counter*cashValue3) / cashValue1
		}
	} else if cashValue3 >= cashValue1 && cashValue3 >= cashValue2 {
		if cashValue1 >= cashValue2 {
			value3Counter = cashWithdrawal / cashValue3
			value1Counter = (cashWithdrawal - value3Counter*cashValue3) / cashValue1
			value2Counter = (cashWithdrawal - value3Counter*cashValue3 - value1Counter*cashValue1) / cashValue2
		} else {
			value3Counter = cashWithdrawal / cashValue3
			value2Counter = (cashWithdrawal - value3Counter*cashValue3) / cashValue2
			value1Counter = (cashWithdrawal - value3Counter*cashValue3 - value2Counter*cashValue2) / cashValue1
		}
	}
	rest = cashWithdrawal - value1Counter*cashValue1 - value2Counter*cashValue2 - value3Counter*cashValue3
	if rest == 0 {
		fmt.Println("Банкомат может выдать вам сумму следующими копюрами:")
		fmt.Println(value1Counter, "купюра(ы) номиналом", cashValue1)
		fmt.Println(value2Counter, "купюра(ы) номиналом", cashValue2)
		fmt.Println(value3Counter, "купюра(ы) номиналом", cashValue3)
	} else {
		fmt.Println("Банкомат не может выдать вам сумму следующими копюрами:", cashValue1, ",", cashValue2, ",", cashValue3)
	}
}
