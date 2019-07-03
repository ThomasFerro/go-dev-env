package golang

import (
	"fmt"
	"os/exec"

	"github.com/go-dev-env/modules"
	"github.com/go-dev-env/triggers"
)

// NewGoRunnerModule Create a new go runner module with the specified trigger and the dir path
func NewGoRunnerModule(path string, t triggers.Trigger) modules.Module {
	cmd := exec.Command("go", "run", fmt.Sprintf("%v/main.go", path))
	return NewGoModule(t, *cmd)
}
