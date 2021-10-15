package Brawlstars

import (
	"encoding/json"
	"log"
)

func ParsePLayerJson(data []byte) Player {
 p := Player{Client: &Client{}}
 log.Fatal(p, data)
 err := json.Unmarshal(data, &p)
 if err !=nil {
 log.Fatal(err)
 }

 return p
}