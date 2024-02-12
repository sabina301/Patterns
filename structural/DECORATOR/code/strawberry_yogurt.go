package code

type strawberryYogurt struct {
	yogurt *yogurt
}

func NewStrawberryYogurt(baseYogurt *yogurt) *strawberryYogurt {
	return &strawberryYogurt{baseYogurt}
}

func (sYogurt *strawberryYogurt) GetPrice() uint8 {
	return 18
}
