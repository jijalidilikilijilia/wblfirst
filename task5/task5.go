package main

import (
	"fmt"
	"time"
)

// 5. Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

func main() {
	ch := make(chan int)

	// Секунды
	N := 5

	// Горутина на запись в канал
	go send(ch)

	//Горутина на чтение из канала
	go receive(ch)

	// Ждем N секунд
	time.Sleep(time.Duration(N) * time.Second)
	fmt.Println("Программа завершена.")
}

func send(ch chan int) {
	//Постоянно отправляем данные на 1 больше прошлого
	for i := 1; ; i++ {
		ch <- i
		fmt.Printf("Отправлено: %d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func receive(ch chan int) {
	//Постоянно читаем данные
	for {
		data := <-ch
		fmt.Printf("Получено: %d\n", data)
	}
}
