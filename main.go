package main

import (
	"crawler/core"
	"fmt"
)

const (
	CONFIG_FILE = "./crawler_config"
)

func main() {
	fmt.Println("Application entry point.")
	cfg, err := core.ParseConfig(CONFIG_FILE)
    if err != nil {
	    fmt.Printf("%v\n", err.Error())
        return
    }

    fmt.Println("Crawler config file successfully read")
    for key, val := range cfg.CfgKeyValues {
        fmt.Printf("%v : %v\n", key, val)
    }
}
