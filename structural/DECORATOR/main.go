package main

import (
	"fmt"
	"patterns/code"
)

// Позволяет добавлять объектам новую функциональность или дополнять объект, не изменяя главный объект
// Здесь - главный объект - это йогурт

func main() {
	yogurt := code.NewYogurt()

	bananaYogurt := code.NewBananaYogurt(yogurt)

	strawberryYogurt := code.NewStrawberryYogurt(yogurt)

	fmt.Println(yogurt.GetPrice())
	fmt.Println(bananaYogurt.GetPrice())
	fmt.Println(strawberryYogurt.GetPrice())
}
