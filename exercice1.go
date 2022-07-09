package main

import (
	"math/rand"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

type Exercise1 struct {
	Rect *sdl.Rect

	Numbers           []int32
	ActiveNumberIndex int32

	Progress
}

func NewExercise1(windowWidth int32, windowHeight int32) *Exercise1 {
	result := Exercise1{
		Rect:              &sdl.Rect{X: 0, Y: 0, W: windowWidth, H: windowHeight},
		Numbers:           make([]int32, 32),
		ActiveNumberIndex: 0,
	}

	result.Progress = NewProgress(result.Rect)

	return &result
}

func (ex *Exercise1) Open() {
	ex.ActiveNumberIndex = 0
	ex.Progress.Update(0.0)
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
	n := ex.Numbers[ex.ActiveNumberIndex]

	if input.AlphaNumbers[n] {
		ex.ActiveNumberIndex += 1
		if ex.ActiveNumberIndex == int32(len(ex.Numbers)) {
			ex.ActiveNumberIndex -= 1

			app.OpenScene("menu")
		}

		ex.Progress.Update(float32(ex.ActiveNumberIndex) / float32(len(ex.Numbers)))
	}
}

func (ex *Exercise1) Render(renderer *sdl.Renderer, app *App) {
	font := app.Fonts["xxxxl"]

	numberAsString := strconv.Itoa(int(ex.Numbers[ex.ActiveNumberIndex]))
	numberWidth := font.GetStringWidth(numberAsString)
	numberRect := sdl.Rect{
		X: ex.Rect.X + ex.Rect.W/2 - numberWidth/2,
		Y: ex.Rect.Y + ex.Rect.H/2 - font.Size/2,
		W: numberWidth,
		H: font.Size,
	}

	ex.Progress.Render(renderer)

	DrawText(renderer, &font, numberAsString, &numberRect, sdl.Color{R: 255, G: 255, B: 255, A: 255})
}
