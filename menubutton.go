package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type MenuButton struct {
	Rect *sdl.Rect
	Name string

	IsHovered bool
	IsActive  bool

	Callback func()
}

func NewMenuButton(containerRect *sdl.Rect, name string, callback func()) (result MenuButton) {
	// Rect.Y does not matter, the container will place the buttons correctly on the y axis
	result.Rect = &sdl.Rect{X: containerRect.X, Y: 0, W: containerRect.W, H: 48}
	result.Name = name

	result.Callback = callback
	return
}

func (btn *MenuButton) Resize(containerRect *sdl.Rect) {
	btn.Rect.X = containerRect.X
}

func (btn *MenuButton) Tick(input *Input) {
	if input.MousePosition.InRect(btn.Rect) {
		btn.IsHovered = true
		btn.IsActive = input.LMB == Pressed

		if input.LMB == JustReleased {
			btn.Callback()
		}
	} else {
		if btn.IsActive {
			btn.IsActive = input.LMB == Pressed
		}

		btn.IsHovered = false
	}
}

func (btn *MenuButton) Render(renderer *sdl.Renderer, app *App) {
	font := app.Fonts["s"]
	nameWidth := font.GetStringWidth(btn.Name)
	nameRect := sdl.Rect{
		X: btn.Rect.X + btn.Rect.W/2 - nameWidth/2,
		Y: btn.Rect.Y + btn.Rect.H/2 - int32(font.Size)/2,
		W: nameWidth,
		H: int32(font.Size),
	}

	if btn.IsActive {
		offsetRect := CloneRect(btn.Rect)
		offsetRect.Y += 2
		nameRect.Y += 2
		DrawRect(renderer, &offsetRect, sdl.Color{R: 100, G: 146, B: 62, A: 255})
		DrawText(renderer, &font, btn.Name, &nameRect, sdl.Color{R: 18, G: 23, B: 31, A: 255})
	} else if btn.IsHovered {
		DrawRect(renderer, btn.Rect, sdl.Color{R: 100, G: 146, B: 62, A: 255})
		DrawText(renderer, &font, btn.Name, &nameRect, sdl.Color{R: 18, G: 23, B: 31, A: 255})
	} else {
		DrawRectOutline(renderer, btn.Rect, sdl.Color{R: 100, G: 146, B: 62, A: 255}, 2)
		DrawText(renderer, &font, btn.Name, &nameRect, sdl.Color{R: 100, G: 146, B: 62, A: 255})
	}
}
