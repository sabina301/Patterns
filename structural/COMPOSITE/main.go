package main

import "patterns/code"

// Применяется когда хотим подогнать объекты под один интерфейс или когда объекты строятся как древовидная структура
// Плюс - простое управление
// Минус - сильная зависимость от этого интерфейса
// Здесь этот интерфейс называется "component"

func main() {
	file1 := code.NewFile("file1")
	file2 := code.NewFile("file2")
	folder1 := code.NewFolder("folder1")

	folder1.AddObject(file1)
	folder1.AddObject(file2)

	file1.Search("cat1")
	file2.Search("cat2")

	folder1.Search("cat3")
}
