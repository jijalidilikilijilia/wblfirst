package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	// Устанавливаем начальные границы поиска
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // Вычисляем середину текущего диапазона

		if arr[mid] == target {
			return mid // Элемент найден, возвращаем его индекс
		} else if arr[mid] < target {
			left = mid + 1 // Ищем в правой половине
		} else {
			right = mid - 1 // Ищем в левой половине
		}
	}

	// Если элемент не найден
	return -1
}

func main() {
	// Создаем сортированный слайс целых чисел
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 3

	// Вызываем нашу функцию бинарного поиска
	index := binarySearch(nums, target)

	// Проверяем, найден ли элемент и выводим результат
	if index != -1 {
		fmt.Printf("Элемент %d найден на %d позиции\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
