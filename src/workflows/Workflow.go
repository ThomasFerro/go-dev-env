package workflows

import "github.com/go-dev-env/triggers"

// Workflow A workflow to be executed in the dev env
type Workflow interface {
	Execute()
	Init()
	Trigger() triggers.Trigger
}
