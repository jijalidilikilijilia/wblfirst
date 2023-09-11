package main

import (
	"fmt"
	"strings"
)

// 20. Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow»

func reverse(words string) string {
	// Разделяем строку на слайс, разделённый пробелом
	str := strings.Fields(words)

	// Цикл работает пока i меньше половины слайса
	for i := 0; i < len(str)/2; i++ {
		// Меняем местами 1 и последний значения и сужаемся
		// ближе к середине
		str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
	}

	// собираем обратно массив строк в строку и возвращаем
	return strings.Join(str, " ")
}

func main() {
	words := "snow dog sun"
	fmt.Printf("reverse: %v\n", reverse(words))
}
