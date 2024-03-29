package main

import "patterns/code"

// Клонируемый объект предоставляет метод клонирования, который возвращает клонированную копию объекта.
// Позволяет копировать объекты, не вдаваясь в подробности их реализации.
// Решаемая проблема: если ты дефолтно копируешь объект, то скопируются не все поля (некоторые приватны).
// Также изза этого происходит привязка к какойто структуре

func main() {
	cat1 := code.NewCat("Musya", 10)
	cat2 := cat1.Clone()
	cat1.ToldAboutU()
	cat2.ToldAboutU()

	dog1 := code.NewDog("Sharik", 5)
	dog2 := dog1.Clone()
	dog1.ToldAboutU()
	dog2.ToldAboutU()
}
