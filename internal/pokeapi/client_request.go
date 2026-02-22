package pokeapi

//use generic function to obtain response
func (c *Client) ClientRequest(url string) (LocationAreaGeneral, error) {
	return genericRequest[LocationAreaGeneral](c, url)
}