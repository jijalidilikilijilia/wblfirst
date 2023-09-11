package main

import (
	"fmt"
	"math/big"
)

// 22. Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20

// calculate принимает два указателя на big.Int и оператор (+, -, *, /) и выполняет операцию
// в зависимости от выбранного оператора. Результат операции возвращается как указатель на новый big.Int
func calculate(a, b *big.Int, operator rune) *big.Int {
	result := new(big.Int)

	switch operator {
	case '+':
		result.Add(a, b)
	case '-':
		result.Sub(a, b)
	case '*':
		result.Mul(a, b)
	case '/':
		result.Div(a, b)
	default:
		fmt.Println("Неправильно выбрана операция")
	}

	return result
}

func main() {
	// Создаём числа a и b больше 1 048 576 (2^20)
	a := new(big.Int)
	b := new(big.Int)

	// Устанавливаем значения чисел a и b в десятичной системе счисления
	a.SetString("6000000", 10)
	b.SetString("3000000", 10)

	// Вывод результатов операций
	fmt.Println("+: ", calculate(a, b, '+'))
	fmt.Println("-: ", calculate(a, b, '-'))
	fmt.Println("*: ", calculate(a, b, '*'))
	fmt.Println("/: ", calculate(a, b, '/'))
}
