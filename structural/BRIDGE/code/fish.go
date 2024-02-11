package code

import "fmt"

type Fish struct {
}

func (f *Fish) Digest() {
	fmt.Println("fish is digested\n")
}
