package code

import "fmt"

type Cat struct {
}

func (cat *Cat) SayMeow() {
	fmt.Println("Meow")
}
