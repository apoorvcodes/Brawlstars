package Brawlstars

import (
	"fmt"

	"github.com/gominima/go-requests"
)

type Client struct {
	token string
	proxy bool
	base  string
}

func BsClient(token string, proxy bool) *Client {
	return &Client{token: token, proxy: proxy, base: "https://api.brawlstars.com/v1"}
}

func (c *Client) GetPlayer(tag string) (*Player, error) {
	data, err := goquests.Get(goquests.Request{
		URL: fmt.Sprintf("%v/players/%v", c.base, tag),
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %v", c.token),
		},
		Data: map[string]interface{}{},
	})
	if err != nil {
		return &Player{}, err
	}
	var p *Player
	fmt.Print(data)
	return p, nil
}

func (c *Client) GetClub(tag string) *Client {
	return c
}
