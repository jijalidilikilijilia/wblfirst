package main

import "fmt"

// 11. Реализовать пересечение двух неупорядоченных множеств.

func intersection(set1, set2 []int) []int {
	// Создаем мапу для хранения элементов первого множества
	set1Map := make(map[int]bool)

	// Создаем слайс для хранения пересечения множеств
	result := []int{}

	// Заполняем мапу элементами первого множества
	for _, num := range set1 {
		set1Map[num] = true
	}

	// Проверяем элементы второго множества и добавляем их в результат,
	// если они есть в первом множестве
	for _, num := range set2 {
		if set1Map[num] {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	fmt.Println("Пересечение:", intersection(set1, set2))
}
