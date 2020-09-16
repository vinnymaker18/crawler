package core

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// An error encountered when trying to parse a document. For now, it's just
// a dumb type, doesn't hold any info about the actual error.
type ParseError struct {
}

func (err *ParseError) Error() string {
	return "Parsing error"
}

// Fetch the raw document at the (remote) address given by link.
func FetchRawDocFromURL(link *url.URL) (*RawDocument, error) {
	resp, err := http.Get(link.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	mimeType := http.DetectContentType(body)
	return &RawDocument{mimeType: mimeType, content: body}, nil
}

// ArticleExtractors extract interesting content from raw documents.
// Different implementations may exist for different file types, data
// sources etc...
type ArticleExtractor interface {
	Parse(RawDocument) (Article, error)
}

// A central registry for various extractors. Main module will
// register all extractors.
var extractors = make(map[int]*ArticleExtractor)

func RegisterExtractor(docType int, extractor *ArticleExtractor) {
	extractors[docType] = extractor
}
