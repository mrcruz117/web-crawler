// normalize_url.go
package main

import (
	"net/url"
	"strings"
)

func normalizeURL(input string) (string, error) {
	parsedURL, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	// Remove scheme (http or https)
	host := parsedURL.Host
	if host == "" {
		host = parsedURL.Path
	}

	// Construct the path without trailing slash
	path := strings.TrimSuffix(parsedURL.Path, "/")
	return host + path, nil
}
