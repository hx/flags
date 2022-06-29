package machines

import (
	"github.com/hx/flags/actions"
	"github.com/hx/flags/states"
)

type DualLimits struct {
	states.State
	UnsafeMinimum int
	SafeMinimum   int
	SafeMaximum   int
	UnsafeMaximum int
}

func NewDualLimits(minimum int, maximum int) *DualLimits {
	return &DualLimits{
		State:         1<<maximum - 1,
		UnsafeMinimum: minimum,
		SafeMinimum:   minimum,
		SafeMaximum:   maximum,
		UnsafeMaximum: maximum,
	}
}

func (d *DualLimits) Get() states.State {
	return d.State
}

func (d *DualLimits) Perform(action actions.Action, isUnsafe bool) {
	var (
		state = action.Apply(d.State)
		count = state.Count()
		min   = d.SafeMinimum
		max   = d.SafeMaximum
	)

	if isUnsafe {
		min = d.UnsafeMinimum
		max = d.UnsafeMaximum
	}

	if count >= min && count <= max {
		d.State = state
		return
	}

	diff := state.Diff(d.State)

	// If every flag changed, we can't do anything.
	if len(diff) == max {
		return
	}

	// Start hunting from around the average of the changed flags. E.g., if flags 4 and 6 were changed, start hunting at
	// flag 5.
	var changedSum int
	for i := range diff {
		changedSum += i
	}
	changedAverage := int(float64(changedSum)/float64(len(diff)) + 0.5) // +0.5 causes rounding to the nearest int

	// Look 1 ahead, then 1 behind, then 2 ahead, 2 behind, 3 ahead, and so on.
	for i := 0; i < 2*max; i++ {
		candidate := changedAverage
		if i%2 == 0 {
			candidate += i / 2
		} else {
			candidate -= i / 2
		}
		if candidate < 0 || candidate >= max {
			continue
		}
		if _, found := diff[candidate]; found {
			continue
		}
		if count > max && state.Get(candidate) {
			state = state.Toggle(candidate)
			count -= 1
		} else if count < min && !state.Get(candidate) {
			state = state.Toggle(candidate)
			count += 1
		}
		if count >= min && count <= max {
			d.State = state
			return
		}
	}
}
