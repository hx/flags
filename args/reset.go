package args

import (
	"errors"
	"github.com/hx/flags/actions"
	"github.com/hx/flags/app"
	"strings"
)

func init() {
	registrar["reset"] = func(args []string, config *app.Config) (info string, err error) {
		if config.StateMachine == nil {
			return "", errors.New("reset must be specified after state machine")
		}
		spec := strings.Join(args, ",") // TODO make this unnecessary
		config.Job(&app.Job{
			Action: actions.SetAll{config.StateMachine.Get()},
			Spec:   spec,
		})
		return "On schedule: " + spec, nil
	}
}
