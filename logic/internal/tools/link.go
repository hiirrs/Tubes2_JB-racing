package scraping

import (
	"fmt"
	"strings"
)

func adjustLink(siteURLPrev string) (siteURL string){
	fmt.Println("masuk adjust link")
	if strings.Contains(siteURLPrev, "/wiki/"){
		if strings.Contains(siteURLPrev, "https://") {
			return siteURLPrev
		}
		return "https://en.wikipedia.org" + siteURLPrev
	} else {
		siteURL := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(siteURLPrev, " ", "_")
		return siteURL
	}
}

