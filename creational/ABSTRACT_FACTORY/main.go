package main

import "patterns/code"

// Идея - создание сразу кучи объектов одной структуры
// Например - есть 2 приюта (firstShelter и secondShelter) с котами и с сабаками. обычно берут либо из 1го приюта либо из 2го
// Здесь фабрика - это интерфейс который реализует функции по созданию объектов. Именно его должны имплементировать наши два приюта (две фабрики)
// В обычной фабрике, фабрика - это функция

func makeShelter(shelterName string) code.AbstractFactory {
	if shelterName == "shelter1" {
		return &code.Shelter1{}
	} else if shelterName == "shelter2" {
		return &code.Shelter2{}
	}
	return nil
}

func main() {

	shelter1 := makeShelter("shelter1")
	shelter2 := makeShelter("shelter2")

	cat1 := shelter1.MakeCat()
	dog1 := shelter1.MakeDog()

	cat2 := shelter2.MakeCat()
	dog2 := shelter2.MakeDog()

	cat1.SayMeow()
	cat2.SayMeow()

	dog1.SayAw()
	dog2.SayAw()
}
