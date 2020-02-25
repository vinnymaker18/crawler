package main

import (
	"fmt"

	"crawler/core"
	"crawler/fetchers"
)

const (
	CONFIG_FILE = "crawler_config"
)

func main() {
	cfg, err := core.ParseConfig(CONFIG_FILE)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	fetcher, err := fetchers.NewTweetFetcher(cfg)
	linkItems, err := fetcher.FetchLinks()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	for _, item := range linkItems {
		fmt.Println(item.PageURL)
	}
}
