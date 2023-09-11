package main

import "fmt"

// 21. Реализовать паттерн «адаптер» на любом примере.

// Паттерн Адаптер (Adapter) предназначен для преобразования
// интерфейса одного класса в интерфейс другого. Благодаря реализации данного
// паттерна мы можем использовать вместе классы с несовместимыми интерфейсами.

// Интерфейс для получения каких-либо старых данных
type OldSystem interface {
	GetOldData() string
}

// Имеем старую систему с данными, которые нам требуется получить
type LegacySystem struct{}

func (l *LegacySystem) GetOldData() string {
	return "Данные из старой системы"
}

// Новый интерфейс, который мы хотим использовать
type NewSystem interface {
	GetData() string
}

// Адаптер для преобразования старого интерфейса в новый
type Adapter struct {
	OldSystem OldSystem
}

// GetData реализует метод интерфейса NewSystem и возвращает данные,
// полученные из старой системы через вызова метода GetOldData() структуре,
// хранящейся в поле OldSystem.
func (a *Adapter) GetData() string {
	return a.OldSystem.GetOldData()
}

func main() {
	// Допустим у нас есть старая система
	legacySystem := &LegacySystem{}

	// Создаем для неё адаптер и передаём, т.к. legacySystem
	// реализует интерфейс OldSystem и имеет метод GetOldData
	adapter := &Adapter{
		OldSystem: legacySystem,
	}

	// Используем новый интерфейс для получения данных из старой системы
	data := adapter.GetData()

	fmt.Println(data)
}
