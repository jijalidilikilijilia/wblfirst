package main

import (
	"fmt"
	"sync"
)

// 18. Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.

// структура счетчика
type Counter struct {
	// Мьютекс для безопастного доступа к счетчику
	mu    *sync.RWMutex
	value int
}

// Увеличение счетчика на 1
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// GetValue - функция для получения текущего значения счетчика
func (c *Counter) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func main() {
	counter := &Counter{}

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Запускаем несколько горутин для инкрементирования счетчика
	for i := 0; i < 1300; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.GetValue())
}
