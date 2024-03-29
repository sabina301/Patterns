package main

import (
	"fmt"
	"patterns/code"
)

// Используется, когда должен существовать только 1 экземпляр структуры
// Примеры использования:
// 1. Экземпляр БД — мы хотим создать только один экземпляр объекта БД и этот экземпляр будет использоваться во всём приложении
// 2. Экземпляр Logger - опять же следует создавать только один экземпляр объекта Logger и он должен использоваться во всём приложении
// Объект-одиночка создаётся при первой инициализации структуры. Обычно в структуре, для которого необходимо создать только один экземпляр, определен метод getInstance()
// При вызове getInstance() в первый раз создаёт, а дальше просто возвращает уже созданный один и тот же экземпляр объекта-одиночки
// Структура, создающая объект-одиночку, должна возвращать один и тот же экземпляр каждый раз, когда несколько горутин пытаются получить доступ к ней

func main() {
	for i := 0; i < 10; i++ {
		go code.GetInstance()
	}
	// Scanln аналогичен Scan, но останавливает сканирование, когда находит символ новой строки,
	// Таким образом, программа будет ожидать ввода символа новой строки и только после этого завершит работу
	fmt.Scanln()
}
