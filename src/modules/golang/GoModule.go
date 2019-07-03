package golang

import (
	"bufio"
	"log"
	"os/exec"

	"github.com/go-dev-env/modules"
	"github.com/go-dev-env/triggers"
)

// GoModule A module responsible for running go code
type GoModule struct {
	trigger triggers.Trigger
	cmd     exec.Cmd
}

func (module GoModule) manageTriggerNotifications(triggerNotification chan bool) {
	for {
		select {
		case <-triggerNotification:
			module.Execute()
		}
	}
}

func (module GoModule) runCommand() error {
	stdout, err := module.cmd.StdoutPipe()
	if err != nil {
		log.Printf("Go runner connecting to the command's standard output: %v", err)
		return err
	}

	if err := module.cmd.Start(); err != nil {
		log.Printf("Go runner command failed: %v", err)
		return err
	}

	in := bufio.NewScanner(stdout)

	for in.Scan() {
		log.Printf("Output from the executed program: %v", in.Text())
	}

	if err := in.Err(); err != nil {
		log.Printf("Go runner command failed: %s", err)
		return err
	}

	log.Printf("Go runner Command execution finished")
	return nil
}

// Init Initialize the go runner module and his trigger
func (module GoModule) Init() {
	log.Println("Initializing a go runner module")

	if module.trigger != nil {
		triggerNotifcation := module.trigger.Init()
		go module.manageTriggerNotifications(triggerNotifcation)
	}

	module.Execute()
}

// Execute execute the go program
func (module GoModule) Execute() {
	log.Println("Executing the go runner module's action")
	module.runCommand()
}

// Trigger The module's trigger
func (module GoModule) Trigger() triggers.Trigger {
	return module.trigger
}

// NewGoModule Create a new go runner module with the specified trigger
func NewGoModule(t triggers.Trigger, c exec.Cmd) modules.Module {
	return GoModule{
		trigger: t,
		cmd:     c,
	}
}
