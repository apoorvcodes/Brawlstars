package Brawlstars;


type Client struct {
	token string
	proxy bool
}

func BsClient(token string, proxy bool) *Client {
	return &Client{token:token, proxy:proxy}
}


func (c *Client) GetPlayer(player string) *Client{
	return c
}

