package code

type Shelter2 struct {
}

func (sh Shelter2) MakeCat() *Cat {
	return &Cat{}
}

func (sh Shelter2) MakeDog() *Dog {
	return &Dog{}
}
