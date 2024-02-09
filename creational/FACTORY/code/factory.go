package code

func GetAnimal(s string) Animal {
	if s == "cat" {
		return NewCat()
	} else if s == "dog" {
		return NewDog()
	} else {
		return nil
	}
}
