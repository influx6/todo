package main

import (
	"log"

	"github.com/influx6/appblueprint/app"
	"github.com/influx6/relay/engine"
)

func main() {

	if err := app.App.Load("./config/app.yaml"); err != nil {
		if app.App.Env == "dev" {
			panic(msg)
		}
	}

	engine.AppSignalInit(app.App)
	log.Printf("Sucessfully booted App@%+s", app.App.EngineAddr())
}
