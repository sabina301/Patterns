package code

type yogurt struct {
}

func NewYogurt() *yogurt {
	return &yogurt{}
}

func (y *yogurt) GetPrice() uint8 {
	return 10
}
