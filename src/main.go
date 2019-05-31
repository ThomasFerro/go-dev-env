package main

import (
	"log"

	"github.com/go-dev-env/triggers"
	"github.com/go-dev-env/modules"
)

func main() {
	log.Printf("Test : Init ?");
	trigger := triggers.NewFileWatcherTrigger()

	module := modules.NewGoRunnerModule(trigger)
	module.Init()

	for {}
}
