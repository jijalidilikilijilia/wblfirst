package main

import (
	"fmt"
	"sync"
)

/*
3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов с использованием конкурентных вычислений.
*/

func sumOfSquares(numbers []int) int {
	//Создаем канал для отправки значений квадратов числа
	resCh := make(chan int, len(numbers))
	var wg sync.WaitGroup

	// Для каждого числа в массиве запускаем горутину и считаем квадрат, не забывая про вейтгруппы
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			squre := n * n
			//отправляем результат в канал
			resCh <- squre
		}(num)
	}

	//При завершении всех горутин закрываем канал чтобы не вызывать deadlock
	go func() {
		wg.Wait()
		close(resCh)
	}()

	//Читаем данные из канала и записваем в sum
	sum := 0
	for v := range resCh {
		sum += v
	}

	return sum
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := sumOfSquares(numbers)
	fmt.Println("Сумма квадратов чисел:", sum)
}
