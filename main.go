package main

import (
	"fmt"
	"github.com/hx/flags/app"
	"github.com/hx/flags/hids"
	"github.com/hx/flags/states"
	"os"
)

func main() {
	stdio := hids.NewStdio()
	machine := states.NewClamp(2)
	machine.State = 0b1111
	a := app.NewApp(app.NewConfig(stdio, stdio, machine))
	server := hids.NewHttpServer("127.0.0.1:1234", "hello")
	a.Input(server)
	a.Output(server)
	a.Output(hids.NewPiGPIO(false, 12, 6, 13, 16))
	err := a.Run()[0]
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
