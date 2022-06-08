package main

import (
	"fmt"
	"log"

	"github.com/Jungle20m/golang-base/cmd/service_01/component"
	"github.com/Jungle20m/golang-base/utils"
)

func main() {
	fmt.Println("Hello World")

	config, err := component.LoadConfig("cmd\\service_01\\config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	logger, err := utils.NewLoggerInstance(
		config.Logger.Format,
		config.Logger.Output,
		config.Logger.Level,
		config.Logger.Folder,
		config.Logger.FileName,
	)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("test log")
}
