package engine

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

func Draw(g *Game, renderer *sdl.Renderer) {

	state := &g.State
	//render raycast
	pov := 60
	angleStep := Radians(pov) / float64(g.Resolution)
	barWidth := float64(g.Width) / float64(g.Resolution)
	firstAngle := state.PlayerAngle - Radians(pov/2)
	for a := 0; a < g.Resolution; a++ {
		ray := getRay(state.PlayerPosition.X, state.PlayerPosition.Y, firstAngle+float64(a)*angleStep, g.Scale, g.State.Map)

		angleOffset := g.State.PlayerAngle - ray.Angle
		if angleOffset < 0 {
			angleOffset += math.Pi * 2
		} else if angleOffset > 2*math.Pi {
			angleOffset -= 2 * math.Pi
		}
		adjustedLength := ray.Length * math.Cos(angleOffset)
		if int32(adjustedLength) <= 0 {
			continue
		}

		wallHeight := int32(g.Scale) * g.Height / int32(adjustedLength)
		if wallHeight > g.Height {
			wallHeight = g.Height
		}
		wallOffset := g.Height/2 - wallHeight/2
		floorHeight := g.Height - wallHeight - wallOffset
		ceilHeight := g.Height - wallHeight - wallOffset

		// Render the ceiling
		ceil := sdl.Rect{
			X: int32(float64(a) * barWidth),
			Y: 0,
			W: int32(barWidth),
			H: ceilHeight + 1,
		}
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.FillRect(&ceil)

		// Render the wall
		scaledHeight := wallHeight * int32(g.State.Map.WallHeight)
		if scaledHeight > g.Height-wallOffset {
			scaledHeight = g.Height - wallOffset
		}
		wall := sdl.Rect{
			X: int32(a * int(barWidth)),
			Y: int32(g.Height - scaledHeight - wallOffset),
			W: int32(barWidth),
			H: scaledHeight + 1,
		}

		shade := float64(g.Scale) / ray.Length * 2
		if shade < 0.6 {
			shade = 0.6
		} else if shade > 1 {
			shade = 1
		}
		if ray.VerticalHit {
			shade -= 0.1
		}
		var r, g, b uint8 = 93, 138, 168

		r, g, b = uint8(float64(r)*shade), uint8(float64(g)*shade), uint8(float64(b)*shade)
		renderer.SetDrawColor(r, g, b, 255)
		renderer.FillRect(&wall)

		// Render the floor
		floor := sdl.Rect{
			X: int32(a * int(barWidth)),
			Y: wallOffset + wallHeight,
			W: int32(barWidth),
			H: floorHeight,
		}
		renderer.SetDrawColor(77, 38, 0, 255)
		renderer.FillRect(&floor)

	}

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
	tan := Tan(angle)
	aTan := ATan(angle)

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
				ray.VerticalHit = true
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
				ray.VerticalHit = false
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
