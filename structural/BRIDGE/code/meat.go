package code

import "fmt"

type Meat struct {
}

func (m *Meat) Digest() {
	fmt.Println("meat is digested\n")
}
