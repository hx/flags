package app

import (
	"github.com/hx/flags/actions"
)

type App struct {
	*Config
}

func NewApp(config *Config) *App {
	if config == nil {
		panic("config must not be nil")
	}
	return &App{config}
}

func (a *App) Run() []error {
	type result struct {
		inputIndex int
		err        error
	}
	initialDiff := a.Config.StateMachine.Get().Diff(0)
	for _, output := range a.Config.Outputs {
		output.Update(initialDiff)
	}
	var (
		actionsChan = make(chan actions.Action)
		resultsChan = make(chan result)
		results     = make([]error, len(a.Config.Inputs))
	)
	go func() {
		for signal := range actionsChan {
			a.handle(signal)
		}
	}()
	for i := range a.Config.Inputs {
		res := result{inputIndex: i}
		input := a.Config.Inputs[i]
		go func() {
			res.err = input.Listen(actionsChan)
			resultsChan <- res
		}()
	}
	for range results {
		result := <-resultsChan
		results[result.inputIndex] = result.err
	}
	close(resultsChan)
	close(actionsChan)
	return results
}

func (a *App) handle(action actions.Action) {
	machine := a.Config.StateMachine
	previous := machine.Get()
	action.Perform(machine)
	diff := machine.Get().Diff(previous)
	for _, output := range a.Config.Outputs {
		output.Update(diff)
	}
}
