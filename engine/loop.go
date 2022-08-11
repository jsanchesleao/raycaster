package engine

import (
	"github.com/veandco/go-sdl2/sdl"
)

var winWidth, winHeight int32 = 800, 600

const winTitle = "Raycaster"

type Game struct {
	Width, Height int32
	WindowTitle   string
	Scale         int
	Update        func(*Game)
	Render        func(*Game, *sdl.Renderer)
	Inputs        GameInputs
	State         GameState
}

type GameInputs struct {
	W, A, S, D            bool
	UP, DOWN, LEFT, RIGHT bool
}

func (g Game) GameLoop() (func(), error) {

	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

	destroyFunc := func() {
		if renderer != nil {
			renderer.Destroy()
		}
		if window != nil {
			window.Destroy()
		}
	}

	window, err = sdl.CreateWindow(g.WindowTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		g.Width,
		g.Height,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return destroyFunc, err
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return destroyFunc, err
	}

	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				if t.Repeat > 0 {
					break
				}
				switch t.Keysym.Sym {
				case 119:
					g.Inputs.W = (t.State == sdl.PRESSED)
				case 97:
					g.Inputs.A = (t.State == sdl.PRESSED)
				case 115:
					g.Inputs.S = (t.State == sdl.PRESSED)
				case 100:
					g.Inputs.D = (t.State == sdl.PRESSED)
				case 1073741906:
					g.Inputs.UP = (t.State == sdl.PRESSED)
				case 1073741904:
					g.Inputs.LEFT = (t.State == sdl.PRESSED)
				case 1073741905:
					g.Inputs.DOWN = (t.State == sdl.PRESSED)
				case 1073741903:
					g.Inputs.RIGHT = (t.State == sdl.PRESSED)
				}
			}
		}

		g.Update(&g)
		g.Render(&g, renderer)

		renderer.Present()
		sdl.Delay(10)

	}

	return destroyFunc, nil
}
