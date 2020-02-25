package core

import (
	"net/url"
)

// A single topic page link retrieved by a link fetcher.
type LinkItem struct {
	PageURL *url.URL
}

// A LinkFetcher instance discovers topic page links from a source
// (e.g., twitter, reddit, wikipedia etc...). FetchLinks() is called
// repeatedly every minute or 2 to fetch new links as they're found
// at source.
type LinkFetcher interface {
	FetchLinks() ([]LinkItem, error)
}

// An article is a piece of text with some possible additional metadata.
type ArticleItem struct {
	// Article text
	text string
}
