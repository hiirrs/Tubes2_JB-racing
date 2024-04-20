package main

import (
	"fmt"
	"context"

	"logic/internal/tools"
	"logic/internal/entities"
)

func main() {
	fmt.Println("Hi! Welcome to JB Racing Wikirace!")
	// method := ""
	// startingWikipage := ""
	// targetWikipage := ""
	// fmt.Print("How do you want to have your Wikirace handle (BFS/IDS)? ")
	// fmt.Scanln(&method)
	// fmt.Print("Enter the starting wikipedia page: ")
	// fmt.Scanln(&startingWikipage)
	// fmt.Print("Enter the target wikipedia page: ")
	// fmt.Scanln(&targetWikipage)
	fmt.Println("We are sorry. We can't handle your request yet.")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	startingWikipageTrial := "https://en.wikipedia.org/wiki/United_States"

	var trialNodes *entities.Node
	trialNodes, _ = scraping.GetWikiNodes(ctx, startingWikipageTrial)
	entities.PrintTree(trialNodes, 0)
}