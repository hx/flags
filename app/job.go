package app

import (
	"fmt"
	"github.com/hx/flags/actions"
)

type Job struct {
	actions.Action
	Spec     string
	IsUnsafe bool
}

type BoundJob struct {
	*Job
	app *App
}

func (j *Job) Bind(app *App) *BoundJob {
	return &BoundJob{j, app}
}

func (b *BoundJob) Run() {
	// TODO: don't assume STDOUT is usable
	fmt.Printf("Running scheduled job: %T%+v\n", b.Action, b.Action)
	b.app.handle(actionRequest{b.Action, nil, b.IsUnsafe})
}
