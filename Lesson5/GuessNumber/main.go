package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var customNumber int
	fmt.Print("Загадайте число от 1 до 10: ")
	fmt.Scan(&customNumber)
	var guessed int
	rand.Seed(time.Now().UnixNano())
	guessed = rand.Intn(10) + 1
	if customNumber == guessed {
		fmt.Println("Вы угадали!")
	} else {
		if customNumber > guessed {
			fmt.Println("Мое число", guessed, "меньше загаданного")
			guessed = rand.Intn(10-guessed) + guessed
			if customNumber == guessed {
				fmt.Println("Вы угадали!")
			} else {
				if customNumber > guessed {
					fmt.Println("Мое число", guessed, "меньше загаданного")
					guessed = rand.Intn(10-guessed) + guessed
					if customNumber == guessed {
						fmt.Println("Вы угадали!")
					} else {
						if customNumber > guessed {
							fmt.Println("Мое число", guessed, "меньше загаданного")
							guessed = rand.Intn(10-guessed) + guessed
							if customNumber == guessed {
								fmt.Println("Вы угадали!")
							} else {
								fmt.Println("Вы не угадали.")
							}
						}
					}
				}
			}
		} else {
			fmt.Println("Мое число", guessed, "больше загаданного")
			guessed = rand.Intn(guessed) + 1
			if customNumber == guessed {
				fmt.Println("Вы угадали!")
			} else {
				if customNumber < guessed {
					fmt.Println("Мое число", guessed, "больше загаданного")
					guessed = rand.Intn(guessed) + 1
					if customNumber == guessed {
						fmt.Println("Вы угадали!")
					} else {
						if customNumber < guessed {
							fmt.Println("Мое число", guessed, "больше загаданного")
							guessed = rand.Intn(guessed) + 1
							if customNumber == guessed {
								fmt.Println("Вы угадали!")
							} else {
								fmt.Println("Вы не угадали.")
							}
						}
					}
				}
			}
		}
	}
}
