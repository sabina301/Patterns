package main

import (
	"fmt"
	"patterns/code"
)

// Инъекция зависимостей позволяет передавать зависимости объекту вместо того чтобы объект сам создавал или управлял ими
// Тип здесь service1 := code.NewService(rep1) мы просто передали ему репозиторий

func main() {
	rep1 := &code.Repository1{}
	rep2 := &code.Repository2{}

	service1 := code.NewService(rep1)
	fmt.Println(service1.GetUserById(1))

	service2 := code.NewService(rep2)
	fmt.Println(service2.GetUserById(2))

}
