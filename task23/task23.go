package main

import (
	"fmt"
)

// 23. Удалить i-ый элемент из слайса.

// Способ 1. Через создание нового слайса без i используя срезы
func removeElement1(slice []int, index int) []int {
	// Сначала обрезается исходный слайс до индекса не включительно и затем,
	// используя оператор ... предоставляем остаток слайса как значения и
	// добавляем в возвращаемый слайс
	return append(slice[:index], slice[index+1:]...)
}

// Способ 2: Изменение существующего слайса через вытеснение
func removeElement2(slice []int, index int) []int {
	// Копируем конец слайса от индекса в начало, перезаписывая сам элемент по искомому
	// нами индексу. Возвращаем слайс без последего элемента, т.к. это копия
	// последенего скопированного

	// 1 2 3 4 5 требуется удалить число 3
	// После копирования: 1 2 4 5 5.
	// После удаления последнего элемента: 1 2 4 5

	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

// Способ 3. Через создание нового слайса без i путём копирования всех элементов
// исходного слайса, итерируясь по нему
func removeElement3(slice []int, index int) []int {
	var newSlice []int
	for i, elem := range slice {
		// Если текущий индекс равен заданному - начинаем заново, не добавляя в новый массив
		if i == index {
			continue
		}
		newSlice = append(newSlice, elem)
	}

	return newSlice
}

func main() {
	// Создаем исходный слайс
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Исходный слайс:", slice)

	slice1 := removeElement1(slice, 0)
	fmt.Println("После 1 способа: ", slice1)

	slice2 := removeElement2(slice, 1)
	fmt.Println("После 2 способа::", slice2)

	slice3 := removeElement3(slice, 2)
	fmt.Println("После 3 способа::", slice3)
}
