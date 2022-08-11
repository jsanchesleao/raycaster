package main

import (
	"raycaster/engine"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	runtime.LockOSThread()

	gameMap := engine.Map{
		Width:      10,
		WallHeight: 3,
		Tiles: []engine.Tile{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 0, 0, 0, 1, 0, 1, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 1, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 1, 0, 1,
			1, 0, 0, 0, 1, 1, 0, 0, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 0, 0, 1,
			1, 0, 1, 0, 0, 0, 0, 1, 0, 1,
			1, 0, 0, 0, 0, 0, 0, 1, 0, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		},
		Start: engine.Position{
			X: 120,
			Y: 80,
		},
	}

	gameState := engine.GameState{
		Map:            &gameMap,
		PlayerPosition: &gameMap.Start,
		PlayerAngle:    engine.Radians(30),
	}

	game := engine.Game{
		Width:       1024,
		Height:      768,
		Scale:       64,
		Resolution:  1024,
		WindowTitle: "rays",
		State:       gameState,
		Update: func(g *engine.Game) {
			engine.MovePlayer(g)
		},
		Render: func(g *engine.Game, r *sdl.Renderer) {
			r.SetDrawColor(0, 0, 0, 255)
			r.Clear()

			engine.Draw(g, r)

		},
	}

	destroy, err := game.GameLoop()
	defer destroy()

	if err != nil {
		panic(err)
	}
}
