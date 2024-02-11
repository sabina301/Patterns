package code

import "fmt"

type Cat struct {
	food Food
}

func NewCat() *Cat {
	return &Cat{}
}

func (c *Cat) Eat() {
	fmt.Println("cat is eating")
	c.food.Digest()
}

func (c *Cat) AddFood(f Food) {
	c.food = f
}
