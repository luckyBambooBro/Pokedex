package pokeapi

// use generic function to obtain response
func (c *Client) RequestGeneralLocationArea(url string) (*LocationAreaGeneral, error) {
	return genericRequest[LocationAreaGeneral](c, url)
}
