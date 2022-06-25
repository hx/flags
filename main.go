package main

import (
	"fmt"
	"github.com/hx/flags/app"
	"github.com/hx/flags/hids"
	"github.com/hx/flags/states"
	"os"
)

func main() {
	hid := hids.NewStdio()
	machine := states.NewClamp(2)
	machine.State = 0b1111
	err := app.NewApp(app.NewConfig(hid, hid, machine)).Run()[0]
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
