package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

/* 2. Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout. */

// Высчитаем квадрат числа и в конце вызвываем wg.Done
func square(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	square := num * num
	// fmt.Printf("%d^2 = %d\n", num, square)
	io.WriteString(os.Stdout, fmt.Sprintf("%d^2 = %d\n", num, square))
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	//Создание WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутины
	for _, num := range numbers {
		wg.Add(1)
		go square(&wg, num)
	}
	//Ожиданиe завершения всех горутин
	wg.Wait()
}
