package main

import (
	"fmt"
	"sync"
)

// 7. Реализовать конкурентную запись данных в map.

func main() {
	// Создаем мапу с типом int для ключа и значения
	// Т.к. map является ссылочным типом, мы используем make для его инициализации
	mapa := make(map[int]int)

	// Используем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Используем мьютекс для синхронизации доступа к map
	var mu sync.Mutex

	// Запускаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			// Блокируем мьютекс для безопасной записи в map
			mu.Lock()

			mapa[val] = val + 100
			//Снимаем блокировку
			mu.Unlock()
		}(i)
	}

	// Ждём завершения горутин
	wg.Wait()

	// Итерация по мапе
	for k, v := range mapa {
		fmt.Printf("key: %d value: %d\n", k, v)
	}
}
