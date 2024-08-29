package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) map[string]int {

	baseDom, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Print(err)
	}

	rawDom, err := url.Parse(rawCurrentURL)

	if err != nil {
		fmt.Print(err)
	}

	if baseDom.Host != rawDom.Host {
		return pages
	}

	normURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Print(err)
	}

	_, exists := pages[normURL]
	if !exists {
		pages[normURL] = 1
	} else {
		pages[normURL]++
		return pages
	}

	bodyHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("crawling %s", rawCurrentURL)

	URLS, err := getURLsFromHTML(bodyHTML, rawBaseURL)
	if err != nil {
		fmt.Print(err)
	}

	for _, URL := range URLS {
		crawlPage(rawBaseURL, URL, pages)
	}

	return pages
}
