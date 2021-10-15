package Brawlstars

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)


func FetchClient(endpoint string, token string) []byte {
	fetch := http.Client{
		Timeout: time.Second * 200, // Timeout after 2 seconds
	}
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
		"Authorization": []string{"Bearer" +token},
	    }

	res, getErr := fetch.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
          return []byte(err.Error())
	}
	return body
}