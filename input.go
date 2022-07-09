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
	Alpha1        bool
	Alpha2        bool
	Alpha3        bool
	Alpha4        bool
	Alpha5        bool
	Alpha6        bool
	Alpha7        bool
	Alpha8        bool
	Alpha9        bool
	Alpha0        bool
	MousePosition sdl.Point
	LMB           ButtonState
	RMB           ButtonState
}

func (input *Input) Clear() {
	input.Alpha1 = false
	input.Alpha2 = false
	input.Alpha3 = false
	input.Alpha4 = false
	input.Alpha5 = false
	input.Alpha6 = false
	input.Alpha7 = false
	input.Alpha8 = false
	input.Alpha9 = false
	input.Alpha0 = false

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
