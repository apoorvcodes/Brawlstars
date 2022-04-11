package Brawlstars

func BaseUrl(query string, proxy bool) string {
	var endpoint string
	if proxy {
		endpoint = `https://api.brawlstars.com/v1` + query

	}
	if !proxy {
		endpoint = `https://api.brawlstars.com/v1` + query

	}
	return endpoint
}
