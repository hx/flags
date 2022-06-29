package actions

import (
	"github.com/hx/flags/states"
)

type Toggle int

func (t Toggle) Apply(state states.State) states.State {
	return state.Set(int(t), !state.Get(int(t)))
}
