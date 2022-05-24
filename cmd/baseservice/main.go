package main

import (
	"fmt"

	"github.com/Jungle20m/golang-base/cmd/baseservice/component"
)

func main() {
	fmt.Println("Service stating...")

	config, err := component.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("config: %+v", config)
}
