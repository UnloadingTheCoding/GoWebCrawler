package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {

	body := `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
	`

	tests := []struct {
		name       string
		htmlBody   string
		rawBaseURL string
		expected   []string
	}{
		{
			name:       "check absolute and relative",
			htmlBody:   body,
			rawBaseURL: "https://blog.boot.dev",
			expected:   []string{"https://https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.htmlBody, tc.rawBaseURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}

}
