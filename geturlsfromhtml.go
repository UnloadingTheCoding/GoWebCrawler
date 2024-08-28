package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, _ := url.Parse(rawBaseURL)
	htmlReader := strings.NewReader(htmlBody)
	htmlTree, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	var paths []string
	var traverseHTMLNodes func(htmlNodes *html.Node)

	traverseHTMLNodes = func(htmlNodes *html.Node) {
		if htmlNodes.Type == html.ElementNode && htmlNodes.Data == "a" {
			for _, a := range htmlNodes.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						fmt.Printf("ERROR!!! %v", err)
						continue
					}
					newURL := baseURL.ResolveReference(href)
					paths = append(paths, newURL.String())
				}

			}
		}
		for child := htmlNodes.FirstChild; child != nil; child = child.NextSibling {
			traverseHTMLNodes(child)
		}
	}

	traverseHTMLNodes(htmlTree)

	return paths, nil

}
