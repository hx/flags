package app

import "github.com/hx/flags/states"

type Output interface {
	Update(diff states.Diff)
}
