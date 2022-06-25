package states

type Clamp struct {
	State
	Minimum int
}

func NewClamp(minimum int) *Clamp {
	var state State
	for i := 0; i < minimum; i++ {
		state = state.Set(i, true)
	}
	return &Clamp{
		Minimum: minimum,
		State:   state,
	}
}

func (c *Clamp) Get() State {
	return c.State
}

func (c *Clamp) Set(state State) {
	if state.Count() >= c.Minimum {
		c.State = state
	}
}

func (c *Clamp) Toggle(index int) {
	c.Set(c.State.Toggle(index))
}
