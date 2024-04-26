package main

import (
	// "context"
	"fmt"
	// "sync"
	"time"

	// "logic/internal/getPath"
	"logic/internal/entities"
	scraping "logic/internal/tools"
	// "logic/internal/entities"
)

func main() {
	fmt.Println("Hi! Welcome to JB Racing Wikirace!")
	// var childLinks []string
	// var wg sync.WaitGroup
	// var mu sync.Mutex // Add a mutex
	// method := ""
	// startingWikipage := ""
	// targetWikipage := ""
	// fmt.Print("How do you want to have your Wikirace handle (BFS/IDS)? ")
	// fmt.Scanln(&method)
	// fmt.Print("Enter the starting wikipedia page: ")
	// fmt.Scanln(&startingWikipage)
	// fmt.Print("Enter the target wikipedia page: ")
	// fmt.Scanln(&targetWikipage)
	// fmt.Println("We are sorry. We can't handle your request yet.")

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	startingWikipageTrial := "https://en.wikipedia.org/wiki/Ariana_Grande"
	// targetPage := "https://en.wikipedia.org/wiki/Billboard_(magazine)"

	// languageCode := scraping.GetLanguageCode(startingWikipageTrial)
	node := scraping.ScrapeToNode(startingWikipageTrial)
	startTime := time.Now()
	entities.PrintTree(node, 0)
	
	// wg.Add(6)
	// go func() {
	// 	scraping.GetWikiLinks(ctx, startingWikipageTrial, languageCode, &wg, &childLinks)
	// 	wg.Done()
	// }()
	// wg.Wait()

	// mu.Lock()
	// for _, link := range childLinks {
	// 	fmt.Println(link)
	// }
	// mu.Unlock()

	// path := getPath.SearchIDS(startingWikipageTrial, targetPage, ctx, languageCode, 6)
	// if path != nil {
	// 	fmt.Println("The target page is found!")
	// 	getPath.PrintPath(path)
	// } else {
	// 	fmt.Println("The target page is not found!")
	// }
	endTime := time.Now()
	fmt.Println("Duration:", endTime.Sub(startTime))

	// var trialNodes *entities.Node
	// trialNodes, _ = scraping.GetWikiNodes(ctx, startingWikipageTrial)
	// entities.PrintTree(trialNodes, 0)

	// res := scraping.Links{}
	// linksArr := make([]string, 2)
	// linksArr[0] = "https://en.wikipedia.org/wiki/KFC"
	// linksArr[1] = "https://en.wikipedia.org/wiki/Joko_Widodo"

	// res := <-scraping.GetAllWikiLinks(ctx, languageCode, linksArr)
	// scraping.PrintLinks(res)

}
