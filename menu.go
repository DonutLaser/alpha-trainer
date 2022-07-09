package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Menu struct {
	Rect          *sdl.Rect
	ButtonSpacing int32

	Buttons []*MenuButton
}

func NewMenu(windowWidth int32, windowHeight int32) Menu {
	var width int32 = 256
	result := Menu{
		Rect:          &sdl.Rect{X: windowWidth/2 - width/2, Y: 0, W: width, H: windowHeight},
		ButtonSpacing: 20,
		Buttons:       make([]*MenuButton, 0),
	}

	btn1 := NewMenuButton(result.Rect, "Exercise 1", func() { fmt.Println("Exercise 1") })
	result.Buttons = append(result.Buttons, &btn1)
	btn2 := NewMenuButton(result.Rect, "Exercise 2", func() { fmt.Println("Exercise 2") })
	result.Buttons = append(result.Buttons, &btn2)
	btn3 := NewMenuButton(result.Rect, "Exercise 3", func() { fmt.Println("Exercise 3") })
	result.Buttons = append(result.Buttons, &btn3)

	buttonCount := int32(len(result.Buttons))
	totalSpacing := (buttonCount - 1) * result.ButtonSpacing

	// All the buttons should be the same size
	startY := windowHeight/2 - (buttonCount*result.Buttons[0].Rect.H+totalSpacing)/2
	for index, btn := range result.Buttons {
		btn.Rect.Y = int32(index)*result.ButtonSpacing + startY + int32(index)*btn.Rect.H
	}

	return result
}

func (menu *Menu) Resize(windowWidth int32, windowHeight int32) {
	menu.Rect.X = windowWidth/2 - menu.Rect.W/2
	menu.Rect.H = windowHeight

	for _, btn := range menu.Buttons {
		btn.Resize(menu.Rect)
	}

	buttonCount := int32(len(menu.Buttons))
	totalSpacing := (buttonCount - 1) * menu.ButtonSpacing

	startY := windowHeight/2 - (buttonCount*menu.Buttons[0].Rect.H+totalSpacing)/2
	for index, btn := range menu.Buttons {
		btn.Rect.Y = int32(index)*menu.ButtonSpacing + startY + int32(index)*btn.Rect.H
	}
}

func (menu *Menu) Tick(input *Input) {
	for _, btn := range menu.Buttons {
		btn.Tick(input)
	}
}

func (menu *Menu) Render(renderer *sdl.Renderer, app *App) {
	for _, btn := range menu.Buttons {
		btn.Render(renderer, app)
	}
}
