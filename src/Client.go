package Brawlstars;


import "github.com/apoorvcodes/Brawlstars/src/Util"

type Client struct {
	token string
	proxy bool
}

func BsClient(token string, proxy bool) *Client {
	return &Client{token:token, proxy:proxy}
}


func (c *Client) GetPlayer(tag string) []byte{
	value  := Utils.BaseUrl("/player", true)
	data := Utils.FetchClient(value, c.token)
	return data
	
}

func (c *Client) GetClub(tag string) *Client{
	return c
}

