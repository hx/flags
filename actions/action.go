package actions

import "github.com/hx/flags/states"

type Action interface {
	Apply(states.State) states.State
}
