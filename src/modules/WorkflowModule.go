package modules

import (
	"log"

	"github.com/go-dev-env/builders"
	"github.com/go-dev-env/triggers"
)

// WorkflowModule A module responsible for running the workflow
type WorkflowModule struct {
	trigger    triggers.Trigger
	builder    builders.Builder
	sourcePath string
}

func (module WorkflowModule) manageTriggerNotifications(triggerNotification chan bool) {
	for {
		select {
		case <-triggerNotification:
			go module.Execute()
		}
	}
}

func (module WorkflowModule) executeWorkflow() error {
	_, err := module.builder.Build(module.sourcePath)
	return err
}

// Init Initialize the workflow module and his trigger
func (module WorkflowModule) Init() {
	log.Println("Initializing a workflow module")

	if module.trigger != nil {
		triggerNotifcation := module.trigger.Init()
		go module.manageTriggerNotifications(triggerNotifcation)
	}

	go module.Execute()
}

// Execute execute the workflow
func (module WorkflowModule) Execute() {
	log.Println("Executing the workflow")
	module.executeWorkflow()
}

// Trigger The module's trigger
func (module WorkflowModule) Trigger() triggers.Trigger {
	return module.trigger
}

// NewWorkflowModule Create a new workflow module with the specified trigger
func NewWorkflowModule(s string, t triggers.Trigger, b builders.Builder) Module {
	return WorkflowModule{
		sourcePath: s,
		trigger:    t,
		builder:    b,
	}
}
