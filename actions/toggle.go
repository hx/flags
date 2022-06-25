package actions

import (
	"github.com/hx/flags/states"
)

type Toggle int

func (t Toggle) Perform(state states.Machine) {
	state.Toggle(int(t))
}
