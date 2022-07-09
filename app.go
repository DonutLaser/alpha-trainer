package main

import "github.com/veandco/go-sdl2/sdl"

type Scene interface {
	Tick(input *Input)
	Render(renderer *sdl.Renderer, app *App)
	Resize(windowWidth int32, windowHeight int32)
}

type App struct {
	Scenes      map[string]Scene
	ActiveScene Scene

	Fonts map[string]Font
}

func NewApp(windowWidth int32, windowHeight int32) (result App) {
	result.Scenes = make(map[string]Scene)

	menuScene := NewMenu(windowWidth, windowHeight)
	result.Scenes["menu"] = &menuScene

	result.ActiveScene = result.Scenes["menu"]

	result.Fonts = make(map[string]Font, 0)
	result.Fonts["s"] = LoadFont("assets/fonts/consolab.ttf", 18)

	return
}

func (app *App) Close() {
	font := app.Fonts["s"]
	font.Unload()
}

func (app *App) Resize(windowWidth int32, windowHeight int32) {
	app.ActiveScene.Resize(windowWidth, windowHeight)
}

func (app *App) Tick(input *Input) {
	app.ActiveScene.Tick(input)
}

func (app *App) Render(renderer *sdl.Renderer) {
	renderer.SetDrawColor(18, 23, 31, 255)
	renderer.Clear()

	app.ActiveScene.Render(renderer, app)

	renderer.Present()
}
