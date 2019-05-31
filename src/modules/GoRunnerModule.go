package modules

import (
	"bufio"
	"log"
	"os/exec"

	"github.com/go-dev-env/triggers"
)

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
	// TODO : Extract the technical complexity
	// TODO : Manage debounce
	// TODO : Manage path
	// TODO : Replace log.Fatal with a clean error management
	log.Println("Executing the go runner module's action")
	cmd := exec.Command("go", "run", "../sandbox/main.go")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	in := bufio.NewScanner(stdout)

	for in.Scan() {
		log.Printf("Output from the executed program: %v", in.Text())
	}
	if err := in.Err(); err != nil {
		log.Printf("Go runner command failed: %s", err)
	}

	log.Printf("Go runner Command execution finished")
}

// Trigger The module's trigger
func (module GoRunnerModule) Trigger() triggers.Trigger {
	return module.trigger
}

// NewGoRunnerModule Create a new go runner module with the specified trigger
func NewGoRunnerModule(t triggers.Trigger) Module {
	return GoRunnerModule{t}
}
