package code

import "fmt"

type Dog struct {
}

func NewDog() Animal {
	fmt.Println("Aw")
	return &Dog{}
}
