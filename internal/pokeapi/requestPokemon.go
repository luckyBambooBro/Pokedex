package pokeapi

//func genericRequest[T pokeStruct](c *Client, url string) (*T, error)
func (cfg *Client) RequestPokemon(url string)(*Pokemon, error) {
	return genericRequest[Pokemon](cfg, url)
}