package main

import (
	"fmt"
	"os"
)

func main() {

	baseURL := os.Args

	if len(baseURL) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(baseURL) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl at %s", baseURL[1])
	}

	bodyHTML, err := getHTML(baseURL[1])
	if err != nil {
		return
	}

	fmt.Printf("%s", bodyHTML)

}
