package actions

import (
	"github.com/hx/flags/states"
)

type Action interface {
	Perform(state states.Machine)
}
