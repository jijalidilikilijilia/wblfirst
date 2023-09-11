package main

import "fmt"

// 12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.

func main() {
	// Определяем последовательность строк
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаём мапу для хранения уникальных ключей и заполняем их пустыми значениями
	set := make(map[string]struct{})

	// Добавляем элементы
	for _, value := range arr {
		set[value] = struct{}{}
	}

	// Теперь set содержит уникальные элементы
	res := []string{}
	for k := range set {
		res = append(res, k)
	}

	fmt.Printf("Множество: %v\n", res)
}
