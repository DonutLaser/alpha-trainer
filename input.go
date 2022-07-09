package main

import "github.com/veandco/go-sdl2/sdl"

type ButtonState uint8

const (
	None ButtonState = iota
	JustPressed
	Pressed
	JustReleased
)

type Input struct {
	AlphaNumbers  []bool
	MousePosition sdl.Point
	LMB           ButtonState
	RMB           ButtonState
}

func NewInput() (result Input) {
	result.AlphaNumbers = make([]bool, 10)
	return
}

func (input *Input) Clear() {
	for i := 0; i < 10; i++ {
		input.AlphaNumbers[i] = false
	}

	if input.LMB == JustReleased {
		input.LMB = None
	} else if input.LMB == JustPressed {
		input.LMB = Pressed
	}

	if input.RMB == JustReleased {
		input.RMB = None
	} else if input.RMB == JustPressed {
		input.RMB = Pressed
	}
}
