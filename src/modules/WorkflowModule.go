package modules

import (
	"log"

	"github.com/go-dev-env/modules/docker"
	"github.com/go-dev-env/triggers"
)

// WorkflowModule A module responsible for running the workflow
type WorkflowModule struct {
	trigger    triggers.Trigger
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
	return docker.Execute("build", "-t", "toto", module.sourcePath)
}

// func (module WorkflowModule) runCommand() error {
// 	stdout, err := module.cmd.StdoutPipe()
// 	if err != nil {
// 		log.Printf("Go runner connecting to the command's standard output: %v", err)
// 		return err
// 	}

// 	if err := module.cmd.Start(); err != nil {
// 		log.Printf("Go runner command failed: %v", err)
// 		return err
// 	}

// 	in := bufio.NewScanner(stdout)

// 	for in.Scan() {
// 		log.Printf("Output from the executed program: %v", in.Text())
// 	}

// 	if err := in.Err(); err != nil {
// 		log.Printf("Go runner command failed: %s", err)
// 		return err
// 	}

// 	log.Printf("Go runner Command execution finished")
// 	return nil
// }

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
func NewWorkflowModule(s string, t triggers.Trigger) Module {
	return WorkflowModule{
		sourcePath: s,
		trigger:    t,
	}
}
