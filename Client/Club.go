package brawlstars


type PartialClub struct {
  Client *Client `json:"-"`
  ID TagID `json:"id"`
  Tag string `json:"tag"`
  Name string `json:"name"`
  Role string `json:"role"`
  BadgeID int `json:"badgeId"`
  BadgeURL string `json:"badgeUrl"`
  Members int `json:"members"`
  Trophies int `json:"trophies"`
  RequiredTrophies int `json:"requiredTrophies"`
  OnlineMembers int `json:"onlineMembers"`
}