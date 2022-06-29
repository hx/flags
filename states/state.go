package states

import "math/bits"

type State uint

func (s State) Get(index int) bool {
	return (int(s)>>index)&1 == 1
}

func (s State) Set(index int, value bool) State {
	if value {
		return s | (1 << index)
	} else {
		return s & ^(1 << index)
	}
}

func (s State) Toggle(index int) State {
	return s.Set(index, !s.Get(index))
}

func (s State) Apply(diff Diff) State {
	for index, value := range diff {
		s = s.Set(index, value)
	}
	return s
}

// Len returns the number of flags represented by the state.
func (s State) Len() int {
	return bits.Len(uint(s))
}

// Count returns the number of flags that are true.
func (s State) Count() (count int) {
	for i, l := 0, s.Len(); i < l; i++ {
		if s.Get(i) {
			count++
		}
	}
	return count
}

func (s State) Diff(previous State) Diff {
	diff := Diff{}
	length := s.Len()
	if a := previous.Len(); a > length {
		length = a
	}
	for i := 0; i < length; i++ {
		var (
			before = previous.Get(i)
			after  = s.Get(i)
		)
		if before != after {
			diff[i] = after
		}
	}
	return diff
}
