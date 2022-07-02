package main

import (
	"fmt"
	"github.com/hx/flags/app"
	"github.com/hx/flags/args"
	"github.com/hx/flags/interfaces"
	"os"
)

func main() {
	config, err := args.Read(os.Args[1:], os.Stdout)
	if err != nil {
		abort(err)
	}

	stdio := interfaces.NewStdio()
	config.Input(stdio).Output(stdio)
	a := app.NewApp(config)

	for _, err := range a.Run() {
		if err != nil {
			abort(err)
		}
	}
}

func abort(reason interface{}) {
	fmt.Fprintln(os.Stderr, reason)
	os.Exit(1)
}
