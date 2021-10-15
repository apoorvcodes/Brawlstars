package Brawlstars;


type PartialClub struct {
  Client *Client `json:"-"`
  ID string `json:"id"`
  Tag string `json:"tag"`
  Name string `json:"name"`
  Role string `json:"role"`
  BadgeID int `json:"badgeId"`
  BadgeURL string `json:"badgeUrl"`
  Members int `json:"members"`
  OnlineMembers int `json:"onlineMembers"`
}


type PlayerClub struct {
  Name string `json:"name"`
  Tag string `json:"tag"`
}