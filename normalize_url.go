package main

import (
	"fmt"
	urlx "net/url"
	"strings"
)

func normalizeURL(url string) (string, error) {
	purl, err := urlx.Parse(url)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	formPath := purl.Host + purl.Path

	formPath = strings.ToLower(formPath)

	formPath = strings.TrimSuffix(formPath, "/")

	return formPath, nil
}
