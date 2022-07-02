package args

import (
	"github.com/hx/flags/app"
)

type Handler func(args []string, config *app.Config) (info string, err error)
