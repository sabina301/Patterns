package code

import "fmt"

type folder struct {
	name     string
	children []Сomponent
}

func NewFolder(name string) *folder {
	return &folder{
		name: name,
	}
}

func (folder *folder) AddObject(f *file) {
	folder.children = append(folder.children, f)
}

func (folder *folder) Search(word string) {
	fmt.Printf("Search word = %s for folder with name %s\n", word, folder.name)
	for _, f := range folder.children {
		f.Search(word)
	}
}
