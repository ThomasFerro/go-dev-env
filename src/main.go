package main

import (
	"os"

	"github.com/go-dev-env/triggers"
	"github.com/go-dev-env/modules"
	"github.com/go-dev-env/builders/docker"
)

func getPath() string {
	path := os.Getenv("SRC_PATH")
	if path == "" {
		path = "/src"
	}
	return path
}

func main() {
	path := getPath()
	
	trigger := triggers.NewFileWatcherTrigger(path)
	builder := docker.NewBuilder()

	module := modules.NewWorkflowModule(path, trigger, builder)
	module.Init()

	for {}
}
