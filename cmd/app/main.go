package main

import (
	"log"

	"github.com/zura-t/go_delivery_system-shops/config"
	"github.com/zura-t/go_delivery_system-shops/internal/app"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config", err)
	}

	app.Run(config)
}
