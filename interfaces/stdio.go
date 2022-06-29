package interfaces

import (
	"bufio"
	"fmt"
	"github.com/hx/flags/actions"
	"github.com/hx/flags/states"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stdio struct {
	Input  io.Reader
	Output io.Writer
	state  states.State
	length int
}

func NewStdio() *Stdio {
	return &Stdio{
		Input:  os.Stdin,
		Output: os.Stdout,
	}
}

var mapBoolToWord = map[bool]string{false: "off", true: "on"}

func (s *Stdio) Update(diff states.Diff) {
	s.state = s.state.Apply(diff)
	if s.length < s.state.Len() {
		s.length = s.state.Len()
	}
	nums := make([]string, s.length)
	for i := range nums {
		if s.state.Get(i) {
			nums[i] = fmt.Sprintf("\033[92m%d\033[0m", i+1)
		} else {
			nums[i] = fmt.Sprintf("%d", i+1)
		}
	}
	fmt.Fprintf(s.Output, "[ %s ] ", strings.Join(nums, " "))
	if len(diff) == 0 {
		fmt.Fprintln(s.Output, "No change")
		return
	}
	parts := make([]string, 0, len(diff))
	for i, v := range diff {
		parts = append(parts, fmt.Sprintf("%d %s", i+1, mapBoolToWord[v]))
	}
	fmt.Fprintf(s.Output, "Changed: %s\n", strings.Join(parts, ", "))
}

func (s *Stdio) Listen(perform actions.Performer) error {
	scanner := bufio.NewScanner(s.Input)
	for scanner.Scan() {
		var (
			text   = scanner.Text()
			unsafe = false
		)
		if strings.HasSuffix(text, "!") {
			unsafe = true
			text = text[:len(text)-1]
		}
		if num, err := strconv.Atoi(text); err == nil {
			if num >= 1 && num <= 64 {
				perform(actions.Toggle(num-1), unsafe)
			} else {
				fmt.Fprintln(s.Output, "Expected a number between 1 and 64")
			}
		} else {
			fmt.Fprintln(s.Output, "Expected numeric input")
		}
	}
	return scanner.Err()
}
