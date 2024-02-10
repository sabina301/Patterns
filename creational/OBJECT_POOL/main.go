package main

import (
	"log"
	"patterns/code"
	"strconv"
)

// Заранее создаётся набор инициализированных и готовых к использованию объектов
// Используй, если:
// 1. Стоимость создания объекта класса высока, а количество таких объектов, которые потребуются в конкретный момент времени, невелико
// 2. Объект в пуле является неизменяемым
// 3. По соображениям производительности
// Пример использования: подключение к БД

func main() {
	connections := make([]code.PoolObject, 0)
	for i := 0; i < 5; i++ {
		connection := &code.Connection{Id: strconv.Itoa(i)}
		connections = append(connections, connection)
	}
	pool, _ := code.InitPool(connections)
	conn, _ := pool.Loan()
	err := pool.Retrieve(conn)
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}

	err = pool.Remove(conn)

}
