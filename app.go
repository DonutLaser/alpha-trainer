package main

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Scene interface {
	Open()
	Resize(windowWidth int32, windowHeight int32)
	Tick(input *Input, app *App)
	Render(renderer *sdl.Renderer, app *App)
}

type App struct {
	Scenes      map[string]Scene
	ActiveScene Scene

	Fonts map[string]Font
}

func NewApp(windowWidth int32, windowHeight int32) *App {
	result := &App{
		Scenes: make(map[string]Scene),
		Fonts:  make(map[string]Font),
	}

	result.Scenes["menu"] = NewMenu(windowWidth, windowHeight)
	result.Scenes["ex1"] = NewExercise1(windowWidth, windowHeight)

	result.ActiveScene = result.Scenes["menu"]

	result.Fonts["s"] = LoadFont("assets/fonts/consolab.ttf", 18)
	result.Fonts["xxxxl"] = LoadFont("assets/fonts/consolab.ttf", 256)

	rand.Seed(time.Now().UnixNano())

	return result
}

func (app *App) Close() {
	font := app.Fonts["s"]
	font.Unload()
}

func (app *App) Resize(windowWidth int32, windowHeight int32) {
	app.Scenes["menu"].Resize(windowWidth, windowHeight)
	app.Scenes["ex1"].Resize(windowWidth, windowHeight)
}

func (app *App) Tick(input *Input) {
	app.ActiveScene.Tick(input, app)
}

func (app *App) Render(renderer *sdl.Renderer) {
	renderer.SetDrawColor(18, 23, 31, 255)
	renderer.Clear()

	app.ActiveScene.Render(renderer, app)

	renderer.Present()
}

func (app *App) OpenScene(scene string) {
	app.ActiveScene = app.Scenes[scene]
	app.ActiveScene.Open()
}
