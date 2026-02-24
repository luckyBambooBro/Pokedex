package pokeapi

// use generic function to obtain response
func (c *Client) RequestLocationArea(url string) (*LocationArea, error) {
	return genericRequest[LocationArea](c, url)
}
