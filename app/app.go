package app

import (
	"fmt"
	"github.com/hx/flags/actions"
	"github.com/robfig/cron/v3"
	"time"
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
	var scheduler *cron.Cron
	if len(a.Jobs) > 0 {
		scheduler = cron.New()
		for _, job := range a.Jobs {
			if _, err := scheduler.AddJob(job.Spec, job.Bind(a)); err != nil {
				panic(err) // TODO
			}
		}
		scheduler.Start()
		// TODO: don't assume STDOUT is usable
		fmt.Printf("Scheduler running in time zone %s\n", time.Now().In(scheduler.Location()).Format("MST"))
	}
	type result struct {
		inputIndex int
		err        error
	}
	initialDiff := a.Config.StateMachine.Get().Diff(0)
	for _, output := range a.Config.Outputs {
		output.Update(initialDiff)
	}
	var (
		actionsChan = make(chan actionRequest)
		resultsChan = make(chan result)
		results     = make([]error, len(a.Config.Inputs))
	)
	go func() {
		for action := range actionsChan {
			a.handle(action)
		}
	}()
	for i := range a.Config.Inputs {
		res := result{inputIndex: i}
		input := a.Config.Inputs[i]
		go func() {
			res.err = input.Listen(func(action actions.Action, isUnsafe bool) chan struct{} {
				done := make(chan struct{})
				actionsChan <- actionRequest{action, done, isUnsafe}
				return done
			})
			resultsChan <- res
		}()
	}
	for range results {
		result := <-resultsChan
		results[result.inputIndex] = result.err
	}
	if scheduler != nil {
		scheduler.Stop()
	}
	close(resultsChan)
	close(actionsChan)
	return results
}

func (a *App) handle(action actionRequest) {
	machine := a.Config.StateMachine
	previous := machine.Get()
	machine.Perform(action.action, action.unsafe)
	diff := machine.Get().Diff(previous)
	for _, output := range a.Config.Outputs {
		output.Update(diff)
	}
	if action.done != nil {
		close(action.done)
	}
}

type actionRequest struct {
	action actions.Action
	done   chan struct{}
	unsafe bool
}
