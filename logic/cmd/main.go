package main

import (
	// "encoding/json"
	// "net/http"
	// "log"

	"context"
	"fmt"
	"time"

	"logic/internal/getPath"
	scraping "logic/internal/tools"
	// "logic/internal/entities"
)

// func main(){
// 	http.HandleFunc("/calculate", calculateHandler)
// 	fmt.Println("Server is running on port 8080...")
// 	http.ListenAndServe(":8080", nil)
// }

func main(){
	fmt.Println("Hi! Welcome to JB Racing Wikirace!")

	// var requestData struct {
	// 	startingWikipageURL string json:"startInput"
	// 	targetWikipageURL string json:"finishInput"
	// 	method string json:"algorithm"
	// }

	// err := json.NewDecoder(r.Body).Decode(&requestData)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	startingWikipageURL := "https://en.wikipedia.org/wiki/Rat_king"
	targetWikipageURL := "https://en.wikipedia.org/wiki/Idola_theatri"
	method := "ids"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	languageCode := scraping.GetLanguageCode(startingWikipageURL)

	pathFound := []string{}
	startTime := time.Now()
	if (method == "ids"){
		pathFound = getPath.SearchIDS(startingWikipageURL, targetWikipageURL, languageCode, ctx)
	}
	// } else if (method == "bfs"){
	// 	path = getPath.BFS(ctx, startingWikipageURL, targetWikipageURL, languageCode)
	// }

	if pathFound != nil {
		fmt.Println("The target page is found!")
		getPath.PrintPath(pathFound)
	} else {
		fmt.Println("The target page is not found!")
	}
	endTime := time.Now()
	fmt.Println("Duration:", endTime.Sub(startTime))

	// siteVisited := 0 // ini nanti kubuatkan fungsinya soon bgtz

	// responseData := struct {
	// 	Path      []string json:"path"
	// 	Duration  string   json:"duration"
	// 	SiteVisited int json:"visitedSites"
	// }{
	// 	Path:      path,
	// 	Duration:  endTime.Sub(startTime).String(),
	// 	SiteVisited: siteVisited,
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(responseData)
}