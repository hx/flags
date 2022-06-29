package interfaces

import (
	"github.com/hx/flags/states"
	"github.com/stianeikeland/go-rpio/v4"
)

type PiGPIO struct {
	inverted bool
	pins     []rpio.Pin
}

func NewPiGPIO(inverted bool, pins ...int) *PiGPIO {
	p := PiGPIO{
		inverted: inverted,
		pins:     make([]rpio.Pin, len(pins)),
	}
	if err := rpio.Open(); err != nil {
		panic("unable to open Raspberry Pi GPIO: " + err.Error())
	}
	for i := range pins {
		pin := rpio.Pin(pins[i])
		pin.Output()
		p.pins[i] = pin
		p.setPin(i, false)
	}
	return &p
}

func (p *PiGPIO) Update(diff states.Diff) {
	l := len(p.pins)
	for i, state := range diff {
		if i < l {
			p.setPin(i, state)
		}
	}
}

func (p *PiGPIO) setPin(index int, state bool) {
	if state == p.inverted {
		p.pins[index].Low()
	} else {
		p.pins[index].High()
	}
}
