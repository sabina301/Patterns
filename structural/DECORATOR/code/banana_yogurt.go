package code

type bananaYogurt struct {
	yogurt *yogurt
}

func NewBananaYogurt(baseYogurt *yogurt) *bananaYogurt {
	return &bananaYogurt{baseYogurt}
}

func (bYogurt *bananaYogurt) GetPrice() uint8 {
	return 20
}
