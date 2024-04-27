package main

import (
	"encoding/json"
	"fmt"
	"log"
	"logic/internal/entities"
	"logic/internal/getPath"
	"net/http"
	"strings"
	"time"
)

type RaceRequest struct {
	StartUrl  string `json:"startUrl"`
	FinishUrl string `json:"finishUrl"`
	Algorithm string `json:"algorithm"`
}

type RaceResult struct {
	Found    bool          `json:"found"`
	Duration time.Duration `json:"duration"`
	Degree   int           `json:"degree"`
	Count    int           `json:"count"`
	Path     string        `json:"path"`
}

func handleRace(w http.ResponseWriter, r *http.Request) {
	// isLoading = true
	decoder := json.NewDecoder(r.Body)
	var request RaceRequest
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	startPage := request.StartUrl
	targetPage := request.FinishUrl
	algorithm := request.Algorithm

	maxDepth := 5
	root := &entities.Node{
		URL:      startPage,
		Children: []*entities.Node{},
		Depth:    0,
	}

	var endNode *entities.Node
	var path []string
	var duration time.Duration

	visited := make(map[string]bool)

	if algorithm == "bfs" {
		var reverse []string
		startTime := time.Now()

		endNode = getPath.BFS(startPage, targetPage, visited)

		endTime := time.Now()
		duration = time.Duration((endTime.Sub(startTime)).Milliseconds())
		// timeProcess = float32(duration.Milliseconds())

		if endNode != nil {
			reverse := getPath.Backtrack(endNode, reverse)
			path = getPath.ReverseArray(reverse)
		}
	} else if algorithm == "ids" {
		var reverse []string
		startTime := time.Now()

		endNode = getPath.IDS(root, targetPage, maxDepth, visited)

		endTime := time.Now()
		duration = time.Duration((endTime.Sub(startTime)).Milliseconds())
		// timeProcess = float32(duration.Milliseconds())

		if endNode != nil {
			reverse := getPath.Backtrack(endNode, reverse)
			path = getPath.ReverseArray(reverse)
		}

	}

	var result RaceResult
	if endNode != nil {
		fmt.Println("The target page is found!")
		result = RaceResult{
			Found:    true,
			Duration: duration,
			Degree:   endNode.Depth,
			Count:    len(visited),
			Path:     strings.Join(path, " -> "),
		}
	} else {
		fmt.Println("The target page is not found!")
		result = RaceResult{
			Found:    false,
			Duration: 0.0,
			Degree:   0,
			Count:    0,
			Path:     "Path Not Found",
		}
	}
	// isLoading = false

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Middleware to add CORS headers
func addCORSHeaders(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // This allows any domain; you may want to restrict it to specific domains
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		handler(w, r)
	}
}

// var isLoading bool

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
