package main

import (
	"crawler/core"
	"fmt"
)

const (
	CONFIG_FILE = "/home/vinay/bin/crawler/config"
)

func main() {
	fmt.Println("Application entry point.")
	_, err := core.ParseConfig(CONFIG_FILE)
	fmt.Printf("%v\n", err.Error())
}
