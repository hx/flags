package app

import "github.com/hx/flags/actions"

type Input interface {
	Listen(perform actions.Performer) error
}
