package main

import (
	"errors"
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func main() {
// 	someFunc()
// }

// Проблемы:
// 1. Использование при создании строки побитовых сдвигов, в нашем
// случае 1 << 10 = 2^10 = 1024. Проще и правильнее будет указывать размер строки в виде числа
// 2. Использование глобальной переменной с непонятным содержимым
// 3. В функции someFunc присваивается срез первых 100 байт строки, но
// если строка содержит многобайтовые символы, то срезирование
// может обрезать символ посреди последовательности байтов, что может привести к ошибкам кодировки

func createHugeString(len int) (string, error) {
	if len <= 0 {
		return "", errors.New("некорректный размер строки")
	}

	var str strings.Builder
	for i := 0; i < len; i++ {
		str.WriteString("H")
	}

	return str.String(), nil
}

func someFunc() (string, error) {
	v, err := createHugeString(1024)
	if err != nil {
		fmt.Println("Ошибка при создании большой строки:", err)
		return "", err
	}

	var justString string
	if len(v) < 100 {
		fmt.Println("Длина полученной строки < 100. Возвращаем что есть")
		// На самом деле не очень нужное присваивание, можно просто вернуть v
		justString = v
		return justString, nil
	}

	justString = v[:100]

	return justString, nil
}

func main() {
	str, err := someFunc()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", str)
}

// Что исправлено:
// 1. Добавлены обработки ошибок
// 2. В случае если при создании большой строки размер меньше 100 мы возвращаем что есть
// 3. Переменная justString теперь не глоабальная и объявлена внутри
// someFunc

//Но лучше сразу использовать функцию createHugeString с размером строки = 100
