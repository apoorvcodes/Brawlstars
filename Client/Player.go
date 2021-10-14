package Brawlstars;

type Player struct {
	Client *Client `json:"-"`
	Name string `json:"name"`
	Tag string `json:"tag"`
	ID string `json:"id"`
	BrawlersUnlocked int `json:"brawlersUnlocked"`
	Victories int `json:"victories"`
	SoloShowdownVictories int `json:"soloShowdownVictories"`
	DuoShowdownVictories int `json:"duoShowdownVictories"`
	TotalExp int `json:"totalExp"`
	ExpFmt string `json:"expFmt"`
	ExpLevel int `json:"expLevel"`
	Trophies int `json:"trophies"`
	HighestTrophies int `json:"highestTrophies"`
	AvatarID int `json:"avatarId"`
	AvatarURL string `json:"avatarUrl"`
	BestTimeAsBigBrawler string `json:"bestTimeAsBigBrawler"`
	BestRoboRumbleTime string `json:"bestRoboRumbleTime"`
	HasSkins bool `json:"hasSkins"`
	Club *PartialClub `json:"club"`
	Brawlers []Brawler `json:"brawlers"`
      }