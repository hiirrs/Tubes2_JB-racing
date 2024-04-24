package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"logic/internal/getPath"
	scraping "logic/internal/tools"
)

func main() {
	http.HandleFunc("/calculate", calculateHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		StartInput  string `json:"startInput"`
		FinishInput string `json:"finishInput"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	startingWikipage := "https://en.wikipedia.org/wiki/" + requestData.StartInput
	finishWikipage := "https://en.wikipedia.org/wiki/" + requestData.FinishInput

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	languageCode := scraping.GetLanguageCode(startingWikipage)

	startTime := time.Now()
	path := getPath.SearchIDS(startingWikipage, finishWikipage, ctx, languageCode)
	endTime := time.Now()

	responseData := struct {
		Path     []string `json:"path"`
		Duration string   `json:"duration"`
	}{
		Path:     path,
		Duration: endTime.Sub(startTime).String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}
