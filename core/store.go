package core

// LinkStore stores links fetched but yet to be processed.
type LinkStore interface {

    // Store new links discovered by fetchers.
	StoreNewLinks([]LinkItem) (bool, error)

    // Return a few (upto the specified limit) unprocessed links 
    // in this database.
    GetUnProcessedLinks(int) ([]LinkItem, error)

    // Mark the links as successfully processed, so next calls to 
    // GetUnProcessed don't return these.
    MarkLinksAsProcessed([]LinkItem) (bool, error)
}

// TODO - ArticleStore and LinkStore may be refactored into a single 
// interface.

// ArticleStore stores the articles extracted from the links.
type ArticleStore interface {

    // Store newly retrieved articles.
    StoreNewArticles([]ArticleItem) (bool, error)
}
