package main

import (
	"context"
	"fmt"
	"time"
)

// 25. Реализовать собственную функцию sleep.

// Блокируем выполнение программы на указанное количество секунд,
// ожидая данных в канале. Она блокирует выполнение программы на указанное время
// и затем завершает выполнение функции.
func mySleep1(seconds int) {
	// Функция After создаёт канал типа Time и через указанное время передаёт
	// туда текущее время после указанного нами
	<-time.After(time.Duration(seconds) * time.Second)
}

// Реализация через context где мы так же блокируем выполнения,
// ожидая ctx.Done()
func mySleep2(seconds int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	defer cancel()

	<-ctx.Done()
}

func main() {
	fmt.Println("start")
	mySleep1(3)
	fmt.Println("Прошли первые 3 секунды")
	mySleep2(3)
	fmt.Println("Прошли вторые 3 секунды")
}
