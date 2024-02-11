package code

import "fmt"

type file struct {
	name string
}

func NewFile(name string) *file {
	return &file{
		name: name,
	}
}

func (file *file) Search(word string) {
	fmt.Printf("Search word = %s for folder with name %s\n", word, file.name)
}
