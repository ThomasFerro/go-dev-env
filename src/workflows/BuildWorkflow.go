package workflows

import (
	"log"

	"github.com/go-dev-env/builders"
	"github.com/go-dev-env/triggers"
)

// BuildWorkflow A workflow based on a builder
type BuildWorkflow struct {
	trigger    triggers.Trigger
	builder    builders.Builder
	sourcePath string
}

func (workflow BuildWorkflow) manageTriggerNotifications(triggerNotification chan bool) {
	for {
		select {
		case <-triggerNotification:
			go workflow.Execute()
		}
	}
}

func (workflow BuildWorkflow) executeWorkflow() error {
	_, err := workflow.builder.Build(workflow.sourcePath)
	return err
}

// Init Initialize the workflow and his trigger
func (workflow BuildWorkflow) Init() {
	log.Println("Initializing a workflow")

	if workflow.trigger != nil {
		triggerNotifcation := workflow.trigger.Init()
		go workflow.manageTriggerNotifications(triggerNotifcation)
	}

	go workflow.Execute()
}

// Execute execute the workflow
func (workflow BuildWorkflow) Execute() {
	log.Println("Executing the workflow")
	workflow.executeWorkflow()
}

// Trigger The workflow's trigger
func (workflow BuildWorkflow) Trigger() triggers.Trigger {
	return workflow.trigger
}

// NewBuildWorkflow Create a new workflow with the specified trigger and builder
func NewBuildWorkflow(s string, t triggers.Trigger, b builders.Builder) Workflow {
	return BuildWorkflow{
		sourcePath: s,
		trigger:    t,
		builder:    b,
	}
}
