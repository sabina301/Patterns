package main

import (
	"fmt"
	"patterns/code"
)

func main() {
	cat := code.GetAnimal("cat")
	dog := code.GetAnimal("dog")
	fmt.Println("OK", cat, dog)
}
