package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// 9. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

func main() {
	inputCh := make(chan int)
	outputCh := make(chan int)

	// Используем вейтгруппы для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Читаем из массива и отправляем в первый канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, num := range arr {
			inputCh <- num
			time.Sleep(500 * time.Millisecond)
		}
		close(inputCh)
	}()

	// Подсчёт квадратов и отправка во второй канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range inputCh {
			result := int(math.Pow(float64(num), 2))
			outputCh <- result
		}
		close(outputCh)
	}()

	// Вывод результатов
	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range outputCh {
			fmt.Println(result)
		}
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()
}
