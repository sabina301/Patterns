package code

type Shelter1 struct {
}

func (sh Shelter1) MakeCat() *Cat {
	return &Cat{}
}

func (sh Shelter1) MakeDog() *Dog {
	return &Dog{}
}
