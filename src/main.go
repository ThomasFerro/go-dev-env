package main

import (
	"github.com/go-dev-env/triggers"
	"github.com/go-dev-env/modules"
	"github.com/go-dev-env/builders/docker"
)

func main() {
	// TODO : Se baser sur la config / les paramètres de la ligne de commande
	path := "."
	trigger := triggers.NewFileWatcherTrigger(path)
	builder := docker.NewBuilder()

	module := modules.NewWorkflowModule(path, trigger, builder)
	module.Init()

	for {}
}
