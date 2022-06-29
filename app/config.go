package app

type Config struct {
	StateMachine StateMachine
	Inputs       []Input
	Outputs      []Output
}

func NewConfig(input Input, output Output, state StateMachine) *Config {
	if input == nil || output == nil || state == nil {
		panic("all arguments must be non-nil")
	}
	return &Config{
		Inputs:       []Input{input},
		Outputs:      []Output{output},
		StateMachine: state,
	}
}

func (c *Config) Input(input Input) *Config {
	if input == nil {
		panic("input must not be nil")
	}
	c.Inputs = append(c.Inputs, input)
	return c
}

func (c *Config) Output(output Output) *Config {
	if output == nil {
		panic("output must not be nil")
	}
	c.Outputs = append(c.Outputs, output)
	return c
}
