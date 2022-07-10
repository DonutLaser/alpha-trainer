package main

import (
	"math/rand"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

type NumberState uint8

const (
	Unknown NumberState = iota
	Incorrect
	Correct
)

type Exercise2 struct {
	Rect          *sdl.Rect
	NumberSpacing int32

	Numbers           []int32
	NumberStates      []NumberState
	ActiveNumberIndex int32

	Progress
}

func NewExercise2(windowWidth int32, windowHeight int32) *Exercise2 {
	result := Exercise2{
		Rect:              &sdl.Rect{X: 0, Y: 0, W: windowWidth, H: windowHeight},
		NumberSpacing:     20,
		Numbers:           make([]int32, 128),
		NumberStates:      make([]NumberState, 128),
		ActiveNumberIndex: 0,
	}

	result.Progress = NewProgress(result.Rect)

	return &result
}

func (ex *Exercise2) Open() {
	ex.ActiveNumberIndex = 0
	ex.Progress.Update(0.0)
	for i := 0; i < 128; i++ {
		ex.Numbers[i] = int32(rand.Intn(10))
		ex.NumberStates[i] = Unknown
	}
}

func (ex *Exercise2) Resize(windowWidth int32, windowHeight int32) {
	ex.Rect.W = windowWidth
	ex.Rect.H = windowHeight

	ex.Progress.Resize(ex.Rect)
}

func (ex *Exercise2) Tick(input *Input, app *App) {

	pressedNumberIndex := -1
	for i := 0; i < len(input.AlphaNumbers); i++ {
		if input.AlphaNumbers[i] {
			pressedNumberIndex = i
			break
		}
	}

	if pressedNumberIndex == -1 {
		return
	}

	n := ex.Numbers[ex.ActiveNumberIndex]
	if int32(pressedNumberIndex) == n {
		ex.NumberStates[ex.ActiveNumberIndex] = Correct
	} else {
		ex.NumberStates[ex.ActiveNumberIndex] = Incorrect
	}

	ex.ActiveNumberIndex += 1
	if ex.ActiveNumberIndex == int32(len(ex.Numbers)) {
		ex.ActiveNumberIndex -= 1
		app.OpenScene("menu")
	}

	ex.Progress.Update(float32(ex.ActiveNumberIndex) / float32(len(ex.Numbers)))
}

func (ex *Exercise2) Render(renderer *sdl.Renderer, app *App) {
	font := app.Fonts["xxl"]

	width := 16*font.CharacterWidth + 15*ex.NumberSpacing
	height := 8*font.Size + 7*ex.NumberSpacing
	startX := ex.Rect.X + ex.Rect.W/2 - width/2
	startY := ex.Rect.Y + ex.Rect.H/2 - height/2
	index := 0

	for i := 0; i < 8; i++ {
		for j := 0; j < 16; j++ {
			numberAsString := strconv.Itoa(int(ex.Numbers[index]))
			numberRect := sdl.Rect{
				X: startX + ex.NumberSpacing*int32(j) + font.CharacterWidth*int32(j),
				Y: startY + ex.NumberSpacing*int32(i) + font.Size*int32(i),
				W: font.CharacterWidth,
				H: font.Size,
			}

			color := sdl.Color{R: 255, G: 255, B: 255, A: 255}
			if ex.NumberStates[index] == Incorrect {
				color = sdl.Color{R: 189, G: 46, B: 46, A: 255}
			} else if ex.NumberStates[index] == Correct {
				color = sdl.Color{R: 100, G: 146, B: 62, A: 255}
			}

			DrawText(renderer, &font, numberAsString, &numberRect, color)

			if index == int(ex.ActiveNumberIndex) {
				activeRect := sdl.Rect{
					X: numberRect.X,
					Y: numberRect.Y + numberRect.H - 5,
					W: numberRect.W,
					H: 3,
				}
				DrawRect(renderer, &activeRect, sdl.Color{R: 255, G: 255, B: 255, A: 255})
			}

			index += 1
		}
	}

	ex.Progress.Render(renderer)
}
