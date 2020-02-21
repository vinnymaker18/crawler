package core

import (
	"time"
)

func discoverNewLinks(fetcher LinkFetcher, store LinkStore) {
	for {
		links, err := fetcher.FetchLinks()
		if err != nil {
			break
		}

		res, err := store.StoreNewLinks(links)
		if err != nil || !res {
			break
		}

		time.Sleep(1 * time.Minute)
	}
}

func Run(fetchers []LinkFetcher, store LinkStore) {
	for _, fetcher := range fetchers {
		go discoverNewLinks(fetcher, store)
	}

}
