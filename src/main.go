package main

import (
	"github.com/go-dev-env/triggers"
	"github.com/go-dev-env/modules"
)

func main() {
	// TODO : Se baser sur la config / les paramètres de la ligne de commande
	// TODO : Modules Docker test / run
	// TODO : Module go test
	path := "../sandbox"
	trigger := triggers.NewFileWatcherTrigger(path)

	module := modules.NewWorkflowModule(path, trigger)
	module.Init()

	for {}
}
