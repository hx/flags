package args

import (
	"fmt"
	"github.com/hx/flags/app"
	"github.com/hx/flags/interfaces"
	"strconv"
	"strings"
)

func init() {
	registrar["gpio"] = func(args []string, config *app.Config) (string, error) {
		var (
			inverted = false
			pins     []int
			pinsStr  []string
		)
		for _, arg := range args {
			if arg == "i" || arg == "I" {
				inverted = true
			} else if num, err := strconv.Atoi(arg); err == nil {
				pins = append(pins, num)
				pinsStr = append(pinsStr, arg)
			} else {
				return "", err
			}
		}
		config.Output(interfaces.NewPiGPIO(inverted, pins...))
		message := fmt.Sprintf("Running on pins %s", strings.Join(pinsStr, ", "))
		if inverted {
			message += " (inverted)"
		}
		return message, nil
	}
}
