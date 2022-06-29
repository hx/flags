package machines

import (
	"github.com/hx/flags/actions"
	"github.com/hx/flags/states"
)

type Clamp struct {
	states.State
	Minimum int
}

func (c *Clamp) Perform(action actions.Action, isUnsafe bool) {
	newState := action.Apply(c.State)
	if newState.Count() >= c.Minimum {
		c.State = newState
	}
}

func NewClamp(minimum int) *Clamp {
	var state states.State
	for i := 0; i < minimum; i++ {
		state = state.Set(i, true)
	}
	return &Clamp{
		Minimum: minimum,
		State:   state,
	}
}

func (c *Clamp) Get() states.State {
	return c.State
}
