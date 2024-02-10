package code

import "fmt"

type Dog struct {
	name string
	age  uint8
}

func NewDog(name string, age uint8) *Dog {
	return &Dog{name: name, age: age}
}

func (dog *Dog) Clone() Animal {
	return &Dog{
		name: dog.name,
		age:  dog.age,
	}
}

func (dog *Dog) ToldAboutU() {
	fmt.Printf("Im dog with name %s with age %d", dog.name, dog.age)
}
