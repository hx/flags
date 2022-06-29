package app

import (
	"github.com/hx/flags/actions"
	"github.com/hx/flags/states"
)

type StateMachine interface {
	Get() states.State
	Perform(action actions.Action, isUnsafe bool)
}
