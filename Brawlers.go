package Brawlstars

type Brawler struct {
	Name            string `json:"name"`
	HasSkin         bool   `json:"hasSkin"`
	Skin            string `json:"skin"`
	Trophies        int    `json:"trophies"`
	HighestTrophies int    `json:"highestTrophies"`
	Power           int    `json:"power"`
	Rank            int    `json:"rank"`
}
