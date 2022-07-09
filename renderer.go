package main

import "github.com/veandco/go-sdl2/sdl"

func DrawRect(renderer *sdl.Renderer, rect *sdl.Rect, color sdl.Color) {
	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	renderer.FillRect(rect)
}

func DrawRectOutline(renderer *sdl.Renderer, rect *sdl.Rect, color sdl.Color, outlineWidth int32) {
	renderer.SetDrawColor(color.R, color.G, color.B, color.A)

	top := sdl.Rect{X: rect.X, Y: rect.Y, W: rect.W, H: outlineWidth}
	right := sdl.Rect{X: rect.X + rect.W - outlineWidth, Y: rect.Y, W: outlineWidth, H: rect.H}
	bottom := sdl.Rect{X: rect.X, Y: rect.Y + rect.H - outlineWidth, W: rect.W, H: outlineWidth}
	left := sdl.Rect{X: rect.X, Y: rect.Y, W: outlineWidth, H: rect.H}

	renderer.FillRect(&top)
	renderer.FillRect(&right)
	renderer.FillRect(&bottom)
	renderer.FillRect(&left)
}

func DrawText(renderer *sdl.Renderer, font *Font, text string, rect *sdl.Rect, color sdl.Color) {
	surface, _ := font.Data.RenderUTF8Blended(text, color)
	defer surface.Free()

	texture, _ := renderer.CreateTextureFromSurface(surface)
	defer texture.Destroy()

	renderer.Copy(texture, nil, rect)
}
