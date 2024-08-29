package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("error: %w", err)
	}

	if res.Header.Get("Content-Type") != "text/html" {
		return "", fmt.Errorf("incorrect content-type")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
