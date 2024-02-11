package code

import "fmt"

type Dog struct {
	food Food
}

func NewDog() *Dog {
	return &Dog{}
}

func (d *Dog) Eat() {
	fmt.Println("dog is eating")
	d.food.Digest()
}

func (d *Dog) AddFood(f Food) {
	d.food = f
}
