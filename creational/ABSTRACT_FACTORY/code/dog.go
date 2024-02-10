package code

import "fmt"

type Dog struct {
}

func (dog *Dog) SayAw() {
	fmt.Println("Aw")
}
