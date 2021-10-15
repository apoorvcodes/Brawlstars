package Brawlstars;

type Player struct {
	Client *Client `json:"-"`
	Name string `json:"name"`
	Tag string `json:"tag"`
	ID string `json:"id"`
	IsQualified bool `json:"isQualifiedFromChampionshipChallenge"`
	TrioVictories int `json:"3vs3Victories"`
	Trophies int `json:"trophies"`
	ExpLevel int `json:"expLevel"`
	ExpPoints int `json:"expPoints"`
	SoloVictories int `json:"soloVictories"`
	DuoVictories int `json:"duoVictories"`
	BestTime int `json:"bestRoboRumbleTime"`
	BestTimeBigBrawler int `json:"bestTimeAsBigBrawler"`
	Club *PlayerClub `json:"club"`
	Brawlers []Brawler `json:"brawlers"`
      }