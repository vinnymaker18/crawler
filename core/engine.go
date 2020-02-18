package core

import (
    "time"
)

func crawl(fetcher Fetcher, store LinkStore) {
    for {
        links, err := fetcher.Fetch()
        if err != nil {
            break
        }

        res, err := store.StoreItems(links)
        if err != nil || !res {
            break
        }

        time.Sleep(1 * time.Minute)
    }
}

func Run(fetchers []Fetcher, store LinkStore) {
    for _, fetcher := range fetchers {
        go crawl(fetcher, store)
    }
}
