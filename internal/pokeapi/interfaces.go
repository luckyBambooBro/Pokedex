package pokeapi

func (LocationAreaGeneral) pokeStructEmptyMethod() {}
func (LocationArea) pokeStructEmptyMethod() {}
func (Pokemon) pokeStructEmptyMethod() {}

type pokeStruct interface {
	pokeStructEmptyMethod()
}