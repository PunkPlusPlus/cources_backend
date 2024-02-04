package main

import (
	"log"

	"github.com/PunkPlusPlus/cources_service/app/config"
	"github.com/PunkPlusPlus/cources_service/app/internal/app"
)

func main() {
	cfg := config.GetConfig()
	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	application.Serve()
}
