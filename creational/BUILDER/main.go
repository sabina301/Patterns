package main

import (
	"fmt"
	"patterns/code"
)

func main() {
	builder := code.NewBuilder()
	cat := builder.Name("Musya").Age(3).Height(20).Build()

	fmt.Printf("%#v\n", cat)

	cat.ToldAboutU()
}
