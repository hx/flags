package main

import (
	"fmt"
	"github.com/hx/flags/app"
	"github.com/hx/flags/interfaces"
	"github.com/hx/flags/machines"
	"os"
)

func main() {
	stdio := interfaces.NewStdio()
	machine := machines.NewDualLimits(2, 4)
	machine.UnsafeMinimum = 1
	server := interfaces.NewHttpServer("127.0.0.1:1234", "hello")
	a := app.NewApp(
		app.NewConfig(stdio, stdio, machine).
			Output(interfaces.NewPiGPIO(false, 21, 20, 16, 12)).
			Input(server).Output(server),
	)
	for _, err := range a.Run() {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
