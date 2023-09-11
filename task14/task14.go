package main

import (
	"fmt"
	"reflect"
)

// 14. Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

type testType struct {
	value int
}

func main() {
	// Вызываем функцию для проверки на разных типах
	checkType(42)
	checkType("string")
	checkType(false)
	checkType(make(chan bool))
	checkType(make(chan int))
	checkType(testType{1})
}

func checkType(val interface{}) {
	// Используем типовое утверждение (type assertion) для определения типа переменной val.
	// Оно используется для проверки и приведения (assertion) интерфейсной переменной к конкретному типу

	switch val.(type) {
	case int:
		fmt.Println("Тип переменной: int")
	case string:
		fmt.Println("Тип переменной: string")
	case bool:
		fmt.Println("Тип переменной: bool")
	case chan int:
		fmt.Println("Тип переменной: chan int")
	default:
		// Если тип неизвестен, используем пакет reflect для вывода фактического типа
		fmt.Printf("Неизвестный тип: %v\n", reflect.TypeOf(val))
	}
}
