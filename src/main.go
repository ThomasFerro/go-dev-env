package main

import (
	"os"

	"github.com/go-dev-env/triggers"
	"github.com/go-dev-env/workflows"
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

	workflow := workflows.NewBuildWorkflow(path, trigger, builder)
	workflow.Init()

	for {}
}
