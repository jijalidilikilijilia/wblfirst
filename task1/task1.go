package main

import "fmt"

/* 1. Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования). */

// Создаем структуру Human с произвольным количеством полей и методов
type Human struct {
	Name string
	Age  int
	City string
}

//Определяем родительский метод
func (h *Human) sayHello() {
	fmt.Printf("I'm %s, I'm %d years old from %s\n", h.Name, h.Age, h.City)
}

//Используем встраивание
type Action struct {
	Human
	Job string
}

// Определяем дочерний метод
func (a *Action) doWork() {
	fmt.Printf("I'm %s and I'm a %s\n", a.Name, a.Job)
}

func main() {
	//Определяем человека
	human := Human{
		Name: "Test",
		Age:  100,
		City: "Stambul",
	}

	action := Action{
		Human: human,
		Job:   "Manager",
	}

	//Получаем доступ к функции sayHello через Human встроенный в Action
	action.sayHello()

	// Вызываем метод doWork который не будет доступен для объекта human
	action.doWork()
}

/*Есть так же второй способ

type Action struct {
	Human human
	Job string
}

В таком случае обращение к полям и методам напрямую будет
недоступно и придётся явно обращаться через .Human.MethodName/Field
*/
