package pokeapi

// use generic function to obtain response
func (c *Client) LocationAreaRequest(url string) (LocationArea, error) {
	return genericRequest[LocationArea](c, url)
}
