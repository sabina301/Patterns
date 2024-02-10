package code

type AbstractFactory interface {
	MakeCat() *Cat
	MakeDog() *Dog
}
