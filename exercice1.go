package main

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Exercise1 struct {
	Rect *sdl.Rect

	Progress

	Numbers []int32
}

func NewExercise1(windowWidth int32, windowHeight int32) *Exercise1 {
	result := Exercise1{
		Rect:    &sdl.Rect{X: 0, Y: 0, W: windowWidth, H: windowHeight},
		Numbers: make([]int32, 32),
	}

	result.Progress = NewProgress(result.Rect)

	return &result
}

func (ex *Exercise1) Open() {
	for i := 0; i < 32; i++ {
		ex.Numbers[i] = int32(rand.Intn(10))
	}
}

func (ex *Exercise1) Resize(windowWidth int32, windowHeight int32) {
	ex.Rect.W = windowWidth
	ex.Rect.H = windowHeight

	ex.Progress.Resize(ex.Rect)
}

func (ex *Exercise1) Tick(input *Input, app *App) {
	if input.Alpha0 {
		ex.Progress.Update(13.0 / 32.0)
	}
}

func (ex *Exercise1) Render(renderer *sdl.Renderer, app *App) {
	ex.Progress.Render(renderer)
}
