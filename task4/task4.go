package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

// 4. Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.

// Воркер принимает контекст для обработки конца программы, имеет свой id, принимает канал откуда
// поступают данные и wg для сигнала о его завершении
func worker(ctx context.Context, id int, inputCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Постоянно считываем значения
	for {
		select {
		// Если получен сигнал отмены - выходим из цикла и прерываем работу воркера
		case <-ctx.Done():
			fmt.Printf("Воркер %d остановился\n", id)
			return
		// Получаем значение из канала и выводим
		case v := <-inputCh:
			fmt.Printf("Воркер %d получил значение %d\n", id, v)
		}
	}
}

func main() {
	// Создаем канал для передачи данных от главного потока к воркерам
	ch := make(chan int)

	// Вводим количество воркеров
	var N int
	fmt.Print("Введите количество воркеров: ")
	fmt.Fscan(os.Stdin, &N)

	if N < 1 {
		log.Fatal("Ошибка. Количество воркеров не может быть меньше 1")
	}

	// Отслеживаем завершение прогрммы пользователем при нажатии CTRL+C
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup

	// Создаём N воркеров
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	// В бесконечном цикле постоянно кидаем значения в канал до тех пор
	// пока не получим сигнал об отмене
	isOver := false
	for !isOver {
		select {
		case <-ctx.Done():
			isOver = true
		default:
			ch <- rand.Intn(1000)
		}
	}

	//Закрываем канал и ждём окончания всех горутин
	close(ch)

	wg.Wait()

	fmt.Println("Конец программы")
}
