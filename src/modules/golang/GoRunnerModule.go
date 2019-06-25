package golang

import (
	"bufio"
	"log"
	"os/exec"

	"github.com/go-dev-env/modules"
	"github.com/go-dev-env/triggers"
)

// TODO : Faire un GoModule avec une commande en paramètre pour pouvoir mutualiser le GoRunner et le GoTester

// GoRunnerModule A module responsible for running go code
type GoRunnerModule struct {
	trigger triggers.Trigger
}

func (module GoRunnerModule) manageTriggerNotifications(triggerNotification chan bool) {
	for {
		select {
		case <-triggerNotification:
			module.Execute()
		}
	}
}

func (module GoRunnerModule) runCommand() error {
	// TODO : Manage path
	cmd := exec.Command("go", "run", "../sandbox/main.go")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Go runner connecting to the command's standard output: %v", err)
		return err
	}

	if err := cmd.Start(); err != nil {
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
func (module GoRunnerModule) Init() {
	log.Println("Initializing a go runner module")

	if module.trigger != nil {
		triggerNotifcation := module.trigger.Init()
		go module.manageTriggerNotifications(triggerNotifcation)
	}

	module.Execute()
}

// Execute execute the go program
func (module GoRunnerModule) Execute() {
	log.Println("Executing the go runner module's action")
	module.runCommand()
}

// Trigger The module's trigger
func (module GoRunnerModule) Trigger() triggers.Trigger {
	return module.trigger
}

// NewGoRunnerModule Create a new go runner module with the specified trigger
func NewGoRunnerModule(t triggers.Trigger) modules.Module {
	return GoRunnerModule{
		trigger: t,
	}
}
