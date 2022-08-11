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
	res := 120
	angleStep := Radians(pov) / float64(res)
	firstAngle := state.PlayerAngle - Radians(pov/2)
	renderer.SetDrawColor(150, 150, 150, 255)
	for a := 0; a < res; a++ {
		ray := getRay(state.PlayerPosition.X, state.PlayerPosition.Y, firstAngle+float64(a)*angleStep, g.Scale, g.State.Map)
		renderer.DrawLine(
			int32(state.PlayerPosition.X),
			int32(state.PlayerPosition.Y),
			int32(ray.X),
			int32(ray.Y),
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

func getRay(rx float64, ry float64, angle float64, scale int, m *Map) Ray {

	ray := Ray{
		Origin: Position{X: rx, Y: ry},
		Angle:  angle,
	}

	hx := rx
	hy := ry
	var deltaHx, deltaHy float64

	vx := rx
	vy := ry
	var deltaVx, deltaVy float64
	maxDof := 10
	dof := 0

	cos := math.Cos(angle)
	sin := math.Sin(angle)
	var tan float64
	var aTan float64
	if cos > 0 && cos < 0.001 {
		tan = sin / 0.001
	} else if cos < 0 && cos > -0.001 {
		tan = sin / -0.001
	} else if cos == 0 {
		tan = 10
	} else {
		tan = sin / cos
	}
	if sin > 0 && sin < 0.001 {
		aTan = cos / 0.001
	} else if sin < 0 && sin > -0.001 {
		aTan = cos / -0.001
	} else if sin == 0 {
		aTan = 10
	} else {
		aTan = cos / sin
	}

	if cos > 0 {
		hx = float64(int(rx/float64(scale))*scale) + float64(scale)
		hy = hy + (rx-hx)*tan
		deltaHx = float64(scale)
		deltaHy = deltaHx * tan * -1
	} else {
		hx = float64(int(rx/float64(scale))*scale) - 0.001
		hy = hy + (rx-hx)*tan
		deltaHx = -float64(scale)
		deltaHy = deltaHx * tan * -1
	}

	if sin < 0 {
		vy = float64(int(ry/float64(scale))*scale) + float64(scale)
		vx = vx + (ry-vy)*aTan
		deltaVy = float64(scale)
		deltaVx = deltaVy * aTan * -1
	} else {
		vy = float64(int(ry/float64(scale))*scale) - 0.001
		vx = vx + (ry-vy)*aTan
		deltaVy = -float64(scale)
		deltaVx = deltaVy * aTan * -1
	}

	for dof < maxDof {
		vLen := math.Sqrt(math.Pow(vx-rx, 2) + math.Pow(vy-ry, 2))
		hLen := math.Sqrt(math.Pow(hx-rx, 2) + math.Pow(hy-ry, 2))
		if vLen == hLen {
			vx += deltaVx
			vy += deltaVy
			hx += deltaHx
			hy += deltaHy
			continue
		} else if vLen < hLen {
			tile := m.RetrieveTile(vx, vy, scale)
			if tile != 0 {
				ray.Detect = tile
				ray.X = vx
				ray.Y = vy
				dof = maxDof
			} else {
				vx += deltaVx
				vy += deltaVy
				dof++
			}
		} else {
			tile := m.RetrieveTile(hx, hy, scale)
			if tile != 0 {
				ray.Detect = tile
				ray.X = hx
				ray.Y = hy
				dof = maxDof
			} else {
				hx += deltaHx
				hy += deltaHy
			}
		}
	}

	ray.Length = math.Sqrt(math.Pow(rx-ray.X, 2) + math.Pow(ry-ray.Y, 2))

	return ray

}
