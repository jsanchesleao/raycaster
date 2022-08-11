package main

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

var winWidth, winHeight int32 = 800, 600

const winTitle = "Raycaster"

func main() {
	runtime.LockOSThread()

	var window *sdl.Window

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		renderer.SetDrawColor(90, 210, 25, 255)
		renderer.Clear()

		renderer.SetDrawColor(0, 0, 255, 255)
		renderer.DrawLine(0, 0, 200, 200)

		renderer.Present()
		sdl.Delay(10)
	}

}
