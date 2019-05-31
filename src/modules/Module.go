package modules

import "github.com/go-dev-env/triggers"

// Module A module of the dev env
type Module interface {
	Execute()
	Init()
	Trigger() triggers.Trigger
}
