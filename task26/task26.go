package main

import (
	"fmt"
	"strings"
)

// 26. Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

// Способ через мапу, т.к. при обращении к несуществующему ключу он создаётся
func isSequenceUnique1(input string) bool {
	// Возводим всё в нижний регистр, т.к. по условию функция проверки должна быть
	// регистронезависимой.
	input = strings.ToLower(input)

	// Создаем мапу для отслеживания уникальных символов.
	// т.к. значения нам не нужны и не хочется тратить лишнюю память, то
	// испольуем тип struct{}. Значений типа nil в мапе быть не можем
	mapa := make(map[rune]struct{})

	// Проходим по каждому символу
	for _, char := range input {
		// Если символ уже встречался, возвращаем false
		if _, exists := mapa[char]; exists {
			return false
		}
		// В противном случае отмечаем символ как встреченный
		mapa[char] = struct{}{}
	}

	// Если не было обнаружено повторяющихся символов, возвращаем true
	return true
}

// Способ через bitmap (битовая карта). МИНУСЫ: Работает только с кодировкой ASCII.
// Для работы с UTF-8 потребуется выделить больше памяти и сделать размер bitmap [65536]
func isSequenceUnique2(input string) bool {
	// Возводим всё в нижний регистр, т.к. по условию функция проверки должна быть
	// регистронезависимой.
	input = strings.ToLower(input)

	// Создаем битовую карту размером 128 бит (ASCII символы)
	var bitmap [128]bool

	// Проходим по каждому символу
	for _, char := range input {
		// Получаем код символа как индекс для битовой карты
		index := int(char)

		// Если бит уже был установлен, значит символ повторяется
		if bitmap[index] {
			return false
		}

		// Устанавливаем бит для данного символа
		bitmap[index] = true
	}

	// Если не было обнаружено повторяющихся символов, возвращаем true
	return true
}

func main() {
	fmt.Println("abcd:", isSequenceUnique1("abcd"))           // true
	fmt.Println("abCdefAaf:", isSequenceUnique1("abCdefAaf")) // false
	fmt.Println("aabcd:", isSequenceUnique1("aabcd"))         // false

	fmt.Println("fgjk:", isSequenceUnique2("fgjk"))   // true
	fmt.Println("jJklm:", isSequenceUnique2("jJklm")) // false
	fmt.Println("abOmv:", isSequenceUnique2("abOmv")) // true
}
