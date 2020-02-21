package crawler

import (
	"crawler/core"
	"fmt"
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

	appKey := cfg.CfgKeyValues["twitter-app-key"]
	appSecret := cfg.CfgKeyValues["twitter-app-secret"]
	fmt.Println(appKey, appSecret)
}
