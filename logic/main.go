package main

import (
	"encoding/json"
	"fmt"
	"log"
	"logic/internal/entities"
	"logic/internal/getPath"
	"net/http"
)

type RaceRequest struct {
	StartUrl  string `json:"startUrl"`
	FinishUrl string `json:"finishUrl"`
	Algorithm string `json:"algorithm"`
}

type RaceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func handleRace(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var request RaceRequest
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	startPage := request.StartUrl
	targetPage := request.FinishUrl
	println(startPage)
	println(targetPage)

	maxDepth := 7
	root := &entities.Node{
		URL:      startPage,
		Children: []*entities.Node{},
		Depth:    0,
	}
	result := getPath.SearchIDSC(root, targetPage, maxDepth)

	var response RaceResponse
	if result != nil {
		fmt.Println("The target page is found!")
		response = RaceResponse{
			Status:  "success",
			Message: "Race calculation completed successfully",
		}
	} else {
		fmt.Println("The target page is not found!")
		response = RaceResponse{
			Status:  "error",
			Message: "Failed to find target page",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Middleware to add CORS headers
func addCORSHeaders(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // This allows any domain; you may want to restrict it to specific domains
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// If it's an OPTIONS request, return just the headers
		if r.Method == "OPTIONS" {
			return
		}

		handler(w, r)
	}
}

func main() {
	http.HandleFunc("/api/race", addCORSHeaders(handleRace))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"logic/internal/getPath"
// 	scraping "logic/internal/tools"
// )

// func main() {
// 	http.HandleFunc("/calculate", calculateHandler)
// 	fmt.Println("Server is running on port 8080...")
// 	http.ListenAndServe(":8080", nil)
// }

// func calculateHandler(w http.ResponseWriter, r *http.Request) {
// 	var requestData struct {
// 		StartInput  string `json:"startInput"`
// 		FinishInput string `json:"finishInput"`
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	start := "https://en.wikipedia.org/wiki/" + requestData.StartInput
// 	finish := "https://en.wikipedia.org/wiki/" + requestData.FinishInput

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	languageCode := scraping.GetLanguageCode(startingWikipage)

// 	startTime := time.Now()
// 	path := getPath.SearchIDS(startingWikipage, finishWikipage, ctx, languageCode, 5)
// 	endTime := time.Now()

// 	responseData := struct {
// 		Path     []string `json:"path"`
// 		Duration string   `json:"duration"`
// 	}{
// 		Path:     path,
// 		Duration: endTime.Sub(startTime).String(),
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(responseData)
// }
