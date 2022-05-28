package main

import "fmt"

func main() {
	var n_students, k_groups, studentNumber int
	fmt.Println("Введите количество студентов на курсе:")
	fmt.Scan(&n_students)
	fmt.Println("Введите количество групп:")
	fmt.Scan(&k_groups)
	fmt.Println("Введите номер студента:")
	fmt.Scan(&studentNumber)

	studentsInGroup := n_students / k_groups
	isMoreStudents := n_students % k_groups
	// Если количество студентов не делится на количество групп, то остаток будет ненулевым
	// тогда мы увеличиваем вместимость группы на 1 студента
	if isMoreStudents > 0 {
		studentsInGroup++
	}
	// Определяем номер группы студента
	groupNumber := studentNumber / studentsInGroup
	// Но если номер студента на цело не делится, то он находится в следующей группе
	// так как вместо в предыдущей ему явно не хватит
	if studentNumber%studentsInGroup > 0 {
		groupNumber++
	}
	fmt.Println("Ваша группа:", groupNumber)
}
