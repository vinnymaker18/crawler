// Utility code package.
package utils

import (
	"net/url"
	"regexp"
)

// NewBool returns a pointer to a bool variable with the given value.
func NewBool(value bool) *bool {
	ret := value
	return &ret
}

// Extract all the urls in a piece of text.
func ExtractURLs(text string) ([]*url.URL, error) {

	// For now, only tweet bodies are considered and only urls
	// starting with 'https://t.co/' are extracted.
	re, err := regexp.Compile(`https:\/\/t\.co\/[\w]+`)
	if err != nil {
		return nil, err
	}

	matches := re.FindAllString(text, -1)
	if matches == nil {
		return make([]*url.URL, 0), nil
	}

	urls := make([]*url.URL, 0, 3)
	for _, match := range matches {
		url, err := url.Parse(match)
		if err == nil {
			urls = append(urls, url)
		}
	}

	return urls, nil
}
