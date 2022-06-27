package actions

type Performer func(action Action) chan struct{}
