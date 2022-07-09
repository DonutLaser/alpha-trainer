package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	FailIfError(err)
	defer sdl.Quit()

	err = ttf.Init()
	FailIfError(err)
	defer ttf.Quit()

	sdl.GLSetAttribute(sdl.GL_FRAMEBUFFER_SRGB_CAPABLE, 1)

	window, err := sdl.CreateWindow("alpha-trainer", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_RESIZABLE)
	FailIfError(err)
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	FailIfError(err)
	defer renderer.Destroy()

	windowWidth, windowHeight := window.GetSize()

	app := NewApp(windowWidth, windowHeight)
	input := Input{}

	running := true
	for running {
		input.Clear()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				input.MousePosition.X = t.X
				input.MousePosition.Y = t.Y
			case *sdl.MouseButtonEvent:
				if t.Button == sdl.BUTTON_LEFT {
					if t.State == sdl.RELEASED {
						input.LMB = JustReleased
					} else if t.State == sdl.PRESSED {
						input.LMB = JustPressed
					}
				} else if t.Button == sdl.BUTTON_RIGHT {
					if t.State == sdl.RELEASED {
						input.RMB = JustReleased
					} else if t.State == sdl.PRESSED {
						input.RMB = JustPressed
					}
				}
			case *sdl.KeyboardEvent:
				keycode := t.Keysym.Sym
				switch keycode {
				case sdl.K_1:
					if t.State != sdl.RELEASED {
						input.Alpha1 = true
					}
				case sdl.K_2:
					if t.State != sdl.RELEASED {
						input.Alpha2 = true
					}
				case sdl.K_3:
					if t.State != sdl.RELEASED {
						input.Alpha3 = true
					}
				case sdl.K_4:
					if t.State != sdl.RELEASED {
						input.Alpha4 = true
					}
				case sdl.K_5:
					if t.State != sdl.RELEASED {
						input.Alpha5 = true
					}
				case sdl.K_6:
					if t.State != sdl.RELEASED {
						input.Alpha6 = true
					}
				case sdl.K_7:
					if t.State != sdl.RELEASED {
						input.Alpha7 = true
					}
				case sdl.K_8:
					if t.State != sdl.RELEASED {
						input.Alpha8 = true
					}
				case sdl.K_9:
					if t.State != sdl.RELEASED {
						input.Alpha9 = true
					}
				case sdl.K_0:
					if t.State != sdl.RELEASED {
						input.Alpha0 = true
					}
				}
			case *sdl.WindowEvent:
				if t.Event == sdl.WINDOWEVENT_RESIZED {
					app.Resize(t.Data1, t.Data2)
				}
			}
		}

		app.Tick(&input)
		app.Render(renderer)
	}

	app.Close()
}
