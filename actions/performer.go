package actions

type Performer func(action Action, isUnsafe bool) chan struct{}
