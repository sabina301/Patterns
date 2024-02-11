package main

import "patterns/code"

// Позволяет отделить абстракцию от ее реализации
// У нас есть 1 большой объект (cat и dog). Делим его на под объекты. То есть щас у нас есть отдельно животные и еда
// Тем самым изменяя еду мы не трогаем животных. Животные могут есть любую еду (которая имплементит интерфейс Food)
// Сам "мост" - это метод AddFood

func main() {
	cat := code.NewCat()
	dog := code.NewDog()

	meat := &code.Meat{}
	fish := &code.Fish{}

	cat.AddFood(meat)
	cat.Eat()

	cat.AddFood(fish)
	cat.Eat()

	dog.AddFood(meat)
	dog.Eat()
}
