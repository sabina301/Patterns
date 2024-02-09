package code

import "fmt"

type Cat struct {
}

func NewCat() Animal {
	fmt.Println("Meow")
	return &Cat{}
}
