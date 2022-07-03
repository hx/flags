package actions

import "github.com/hx/flags/states"

type SetAll struct {
	states.State
}

func (s SetAll) Apply(state states.State) states.State {
	return s.State
}
