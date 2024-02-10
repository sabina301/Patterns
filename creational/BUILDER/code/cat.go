package code

import "fmt"

type Cat struct {
	name   string
	age    uint8
	height uint8
}

func (cat *Cat) ToldAboutU() {
	fmt.Printf("Im cat with name %s with age %d and with height %d", cat.name, cat.age, cat.height)
}
