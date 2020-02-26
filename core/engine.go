package core

import (
	"time"
)

func discoverNewLinks(fetcher LinkFetcher, queue JobQueue) {
	for {
		links, err := fetcher.FetchLinks()
		if err != nil {
			break
		}

		urls := make([]string, 0, len(links))
		for i, link := range links {
			urls[i] = link.PageURL.String()
		}

		res, err := queue.Enqueue(urls)
		if err != nil || !res {
			break
		}

		time.Sleep(1 * time.Minute)
	}
}

func Run(fetchers []LinkFetcher, linkQueue JobQueue) {
	for _, fetcher := range fetchers {
		go discoverNewLinks(fetcher, linkQueue)
	}

}
