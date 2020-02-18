package core

type LinkStore interface {
    // Store new links discovered by fetchers.
	StoreNewLinks([]LinkItem) (bool, error)

    // Return a few (upto the specified limit) unprocessed links 
    // in this database.
    GetUnProcessedLinks(int) ([]LinkItem, error)
}
