package engine

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func Draw(g *Game, renderer *sdl.Renderer) {

	state := &g.State
	// draw grid
	for index, tile := range state.Map.Tiles {

		x := (index % state.Map.Width) * g.Scale
		y := (index / state.Map.Width) * g.Scale

		if tile == 0 {
			renderer.SetDrawColor(255, 255, 255, 255)
		} else {
			renderer.SetDrawColor(255, 0, 0, 255)
		}
		tileRect := sdl.Rect{X: int32(x) + 1, Y: int32(y) + 1, W: int32(g.Scale - 2), H: int32(g.Scale - 2)}
		renderer.FillRect(&tileRect)
	}

	//render rays
	pov := 60
	angleStep := Radians(1)
	firstAngle := state.PlayerAngle - Radians(pov/2)
	renderer.SetDrawColor(150, 150, 150, 255)
	for a := 0; a < pov; a++ {
		ray := getRay(state.PlayerPosition.X, state.PlayerPosition.Y, firstAngle+float64(a)*angleStep, g.Scale, g.State.Map)
		dx := math.Cos(ray.Angle) * ray.Length
		dy := -math.Sin(ray.Angle) * ray.Length
		renderer.DrawLine(
			int32(state.PlayerPosition.X),
			int32(state.PlayerPosition.Y),
			int32(state.PlayerPosition.X+dx),
			int32(state.PlayerPosition.Y+dy),
		)
	}

	//draw player
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(&sdl.Rect{
		X: int32(state.PlayerPosition.X) - 3,
		Y: int32(state.PlayerPosition.Y) - 3,
		W: 6,
		H: 6,
	})

	dx := 10 * math.Cos(float64(state.PlayerAngle))
	dy := 10 * math.Sin(float64(state.PlayerAngle)) * -1
	renderer.DrawLine(
		int32(state.PlayerPosition.X),
		int32(state.PlayerPosition.Y),
		int32(state.PlayerPosition.X+dx),
		int32(state.PlayerPosition.Y+dy),
	)

}

// TODO Use smarter algorithm here
func getRay(rx float64, ry float64, angle float64, scale int, m *Map) Ray {
	var length float64 = 0
	var maxLength float64 = 500
	var detect Tile

	var ns float64 = 1
	dx := ns * math.Cos(angle)
	dy := ns * math.Sin(angle) * -1
	for length < maxLength {
		length += ns
		tile := m.RetrieveTile(rx+(length*dx), ry+(length*dy), scale)
		if tile != 0 {
			detect = tile
			break
		}
	}

	return Ray{
		Origin: Position{X: rx, Y: ry},
		Angle:  angle,
		Length: length,
		Detect: detect,
	}
}
