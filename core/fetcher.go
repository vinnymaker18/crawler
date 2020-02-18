package core

import (
	"net/url"
)

// Fetcher interface and its input/output data.

// A single link/item fetched by fetcher.
type LinkItem struct {
	pageURL url.URL
}

// A Fetcher instance fetches links/items from a source.
type Fetcher interface {
	Fetch() ([]LinkItem, error)
}
