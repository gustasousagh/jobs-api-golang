package main

import (
	"fmt"

	"github.com/gustasousagh/jobs-api-golang/config"
	"github.com/gustasousagh/jobs-api-golang/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.ErrorF("Config: %v", err)
		fmt.Println(err)
		return
	}
	router.Initialize()
}
