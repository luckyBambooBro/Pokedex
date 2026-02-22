package pokeapi

func (LocationAreaGeneral) pokeStructEmptyMethod() {}
func (LocationArea) pokeStructEmptyMethod() {}

type pokeStruct interface {
	pokeStructEmptyMethod()
}