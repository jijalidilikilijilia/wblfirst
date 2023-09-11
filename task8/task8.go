package main

import (
	"fmt"
)

// 8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в
// 1 или 0

func main() {
	num := int64(12345)
	fmt.Printf("Число до: %v\n", num)
	setBit(&num, 10, 1)
	fmt.Printf("Число после: %v\n", num)
}

func setBit(num *int64, index int, oneOrTwo int) {
	if index > 63 || index < 0 {
		fmt.Println("Ошибка ввода номера бита. Номер должен быть в диапазоне от 0 до 63.")
		return
	}
	if oneOrTwo > 1 || oneOrTwo < 0 {
		fmt.Println("Ошибка. Устанавливать можно либо на 1, либо на 0")
		return
	}
	if oneOrTwo == 1 {
		*num = *num | (1 << index)
	} else {
		*num = *num | (0 << index)
	}
}
