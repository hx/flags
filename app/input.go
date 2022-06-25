package app

import "github.com/hx/flags/actions"

type Input interface {
	Listen(actions chan actions.Action) error
}
