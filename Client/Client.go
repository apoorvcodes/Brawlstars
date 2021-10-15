package Brawlstars



type Client struct {
	token string
	proxy bool
}

func BsClient(token string, proxy bool) *Client {
	return &Client{token:token, proxy:proxy}
}


func (c *Client) GetPlayer(tag string) Player{
	value  := BaseUrl("/player/" + tag, true)
	data := FetchClient(value, c.token)
	result := ParsePLayerJson(data)
	return result
	
}

func (c *Client) GetClub(tag string) *Client{
	return c
}

