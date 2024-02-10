package code

import "fmt"

type Cat struct {
	name string
	age  uint8
}

func NewCat(name string, age uint8) *Cat {
	return &Cat{name: name, age: age}
}

func (cat *Cat) Clone() Animal {
	return &Cat{
		name: cat.name,
		age:  cat.age,
	}
}

func (cat *Cat) ToldAboutU() {
	fmt.Printf("Im cat with name %s with age %d", cat.name, cat.age)
}
