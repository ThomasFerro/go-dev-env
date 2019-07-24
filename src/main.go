package main

import (
	"os"

	"github.com/go-dev-env/triggers"
	"github.com/go-dev-env/workflows"
	"github.com/go-dev-env/builders/docker"
)

func getFromEnvOrDefault(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		value = defaultValue
	}
	return value
}

func getPath() string {
	return getFromEnvOrDefault("SRC_PATH", "/src")
}

func getArtifactName() string {
	return getFromEnvOrDefault("ARTIFACT_NAME", "my-go-app")
}

func main() {
	path := getPath()
	
	trigger := triggers.NewFileWatcherTrigger(path)
	builder := docker.NewBuilder(getArtifactName())

	workflow := workflows.NewBuildWorkflow(path, trigger, builder)
	workflow.Init()

	for {}
}
