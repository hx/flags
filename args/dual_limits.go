package args

import (
	"errors"
	"fmt"
	"github.com/hx/flags/app"
	"github.com/hx/flags/machines"
	"regexp"
	"strconv"
)

func init() {
	registrar["dual-limits"] = func(args []string, config *app.Config) (string, error) {
		if len(args) != 2 {
			return "", errors.New("expected exactly 2 arguments")
		}
		pattern := regexp.MustCompile(`^(\d+)(?:/(\d+))?$`)
		nums := make([][]int, 2)
		for i, arg := range args {
			var seq []int
			caps := pattern.FindStringSubmatch(arg)
			if caps == nil {
				return "", errors.New("arguments must match " + pattern.String())
			}
			for j, str := range caps {
				if j == 0 || str == "" {
					continue
				}
				num, _ := strconv.Atoi(str)
				seq = append(seq, num)
			}
			if len(seq) == 2 && seq[0] > seq[1] {
				return "", errors.New("arguments should be given in ascending sequence")
			}
			nums[i] = seq
		}
		var (
			min = nums[0]
			max = nums[1]
		)
		machine := machines.NewDualLimits(min[0], max[len(max)-1])
		if len(min) == 2 {
			machine.SafeMinimum = min[1]
		}
		if len(max) == 2 {
			machine.SafeMaximum = max[0]
		}
		config.StateMachine = machine
		message := fmt.Sprintf("Min %d", machine.SafeMinimum)
		if machine.UnsafeMinimum != machine.SafeMinimum {
			message += fmt.Sprintf(" (unsafe %d)", machine.UnsafeMinimum)
		}
		message += fmt.Sprintf(", max %d", machine.SafeMaximum)
		if machine.UnsafeMaximum != machine.SafeMaximum {
			message += fmt.Sprintf(" (unsafe %d)", machine.UnsafeMaximum)
		}
		return message, nil
	}
}
