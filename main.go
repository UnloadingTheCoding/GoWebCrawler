package main

import (
	"fmt"
	"os"
)

func main() {

	baseURL := os.Args
	rawBaseURL := baseURL[1]
	if len(baseURL) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(baseURL) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl at %s", rawBaseURL)
	}

	pages := make(map[string]int)
	pagesComplete := crawlPage(rawBaseURL, rawBaseURL, pages)

	for key, val := range pagesComplete {
		fmt.Println("%s: %d", key, val)
	}
}
