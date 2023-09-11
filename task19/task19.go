package main

import (
	"fmt"
	"unicode/utf8"
)

// 19. Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverse(word string) string {
	// Получаем количество символов
	length := utf8.RuneCountInString(word)

	// Т.к. у нас могут быть unicode символы - используем тип rune
	res := []rune(word)

	// Используем цикл, чтобы перевернуть элементы в срезе
	for i := 0; i < length/2; i++ {
		// Обмениваем символы между началом и концом строки
		res[i], res[length-1-i] = res[length-1-i], res[i]
	}
	return string(res)
}

func main() {
	fmt.Println(reverse("гларыба"))
}
