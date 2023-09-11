package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// 6. Реализовать все возможные способы остановки выполнения горутины.

func main() {
	// Способ 1: Использование канала (Channel)
	done := make(chan bool)
	go func() {
		time.Sleep(3 * time.Second)
		// Отправляем сигнал
		done <- true
	}()
	// Ждем, пока горутина завершится
	<-done
	fmt.Println("Горутина 1 завершена")

	// Способ 2: Использование контекста (Context)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		// Отменяем контекст, чтобы завершить горутину
		cancel()
	}()

	// Ждем отмену контекста
	<-ctx.Done()
	fmt.Println("Горутина 2 завершена")

	// Способ 3: Использование сигналов ОС (Ctrl+C или SIGTERM)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Выполняем какую-то работу в горутине
		time.Sleep(3 * time.Second)
		// Отправляем сигнал завершения в канал
		c <- os.Interrupt
	}()

	// Ждём сигнал завершения
	<-c
	fmt.Println("Горутина 3 завершена")

	// Способ 4: Использование runtime.Goexit()
	go func() {
		time.Sleep(3 * time.Second)
		// Завершаем горутину
		runtime.Goexit()
	}()

	// Ждем, пока горутина завершится
	fmt.Println("Ожидание завершения горутины 4...")
}
