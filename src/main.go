package main

import (
	"github.com/go-dev-env/triggers"
	"github.com/go-dev-env/modules/golang"
)

func main() {
	// TODO : Se baser sur la config / les paramètres de la ligne de commande
	// TODO : Module Docker
	trigger := triggers.NewFileWatcherTrigger()

	module := golang.NewGoRunnerModule(trigger)
	module.Init()

	for {}
}
