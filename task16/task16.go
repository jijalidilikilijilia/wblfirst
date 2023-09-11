package main

import (
	"fmt"
)

// 16. Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка.

func quicksort(arr []int) []int {
	// Eсли массив содержит 1 элемент или меньше, он уже отсортирован
	if len(arr) <= 1 {
		return arr
	}

	//  Выбираем опорный элемент (в данном случае первый элемент)
	pivot := arr[0]

	// Создаем два массива для элементов меньше и больше опорного
	var less, greater []int

	// Разделяем массив на две части
	for _, value := range arr[1:] {
		if value <= pivot {
			// Элементы меньше опорного добавляем в less
			less = append(less, value)
		} else {
			// Элементы больше опорного добавляем в greater
			greater = append(greater, value)
		}
	}

	// Рекурсивно сортируем оба подмассива
	less = quicksort(less)
	greater = quicksort(greater)

	// Объединяем отсортированные подмассивы и опорный элемент
	// Сначала добавляем в слайс с меньшими элементами опорное элемент, т.к. он
	// больше всех значений less и затем добавляем в новый слайс элементы из массива greater
	return append(append(less, pivot), greater...)
}

func main() {
	// Пример использования
	arr := []int{13, 5, 15, 9, 4, 20, 11, 6, 9}
	sortedArr := quicksort(arr)
	fmt.Println(sortedArr)
}
