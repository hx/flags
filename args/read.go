package args

import (
	"errors"
	"fmt"
	"github.com/hx/flags/app"
	"io"
	"regexp"
	"strings"
)

var commandPattern = regexp.MustCompile(`^(\w[-\w]*)\[(.*)]$`)

func Read(commands []string, logFile io.Writer) (*app.Config, error) {
	config := new(app.Config)
	for _, commandStr := range commands {
		match := commandPattern.FindStringSubmatch(commandStr)
		if match == nil {
			return nil, errors.New("invalid command format")
		}
		commandName := match[1]
		command := registrar[commandName]
		if command == nil {
			return nil, errors.New("unknown command")
		}
		lastMachine := config.StateMachine
		info, err := command(strings.Split(match[2], ","), config)
		if lastMachine != nil && config.StateMachine != lastMachine {
			return nil, fmt.Errorf("%s: a state machine has already been specified", commandName)
		}
		if err != nil {
			return nil, fmt.Errorf("%s: %s", commandName, err)
		}
		if info != "" {
			fmt.Fprintf(logFile, "[ %s ] %s\n", commandName, info)
		}
	}
	if config.StateMachine == nil {
		return nil, errors.New("no state machine specified")
	}
	return config, nil
}
