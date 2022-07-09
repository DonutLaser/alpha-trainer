package main

import "github.com/veandco/go-sdl2/sdl"

type Progress struct {
	Rect         *sdl.Rect
	PercentValue float32
	MaxWidth     int32
}

func NewProgress(containerRect *sdl.Rect) (result Progress) {
	result.Rect = &sdl.Rect{X: containerRect.X, Y: containerRect.Y + containerRect.H - 18, W: 0, H: 18}
	result.MaxWidth = containerRect.W

	return
}

func (progress *Progress) Resize(containerRect *sdl.Rect) {
	progress.MaxWidth = containerRect.W
	progress.Rect.X = containerRect.X
	progress.Rect.Y = containerRect.Y + containerRect.H - 18
	progress.Rect.W = int32(float32(progress.MaxWidth) * progress.PercentValue)
}

func (progress *Progress) Tick(input *Input) {
	// There's nothing to do here
}

func (progress *Progress) Render(renderer *sdl.Renderer) {
	DrawRect(renderer, progress.Rect, sdl.Color{R: 100, G: 146, B: 62, A: 255})
}

func (progress *Progress) Update(percent float32) {
	progress.PercentValue = percent
	progress.Rect.W = int32(float32(progress.MaxWidth) * percent)
}
