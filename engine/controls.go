package engine

import "math"

const baseSpeed = 0.03
const baseAngleSpeed = 2
const runSpeed = 0.07

func MovePlayer(g *Game) {
	state := &g.State
	if g.Inputs.W || g.Inputs.UP {
		speed := baseSpeed
		if g.Inputs.SHIFT {
			speed = runSpeed
		}
		px := state.PlayerPosition.X
		py := state.PlayerPosition.Y
		scale := float64(g.Scale)
		dx := math.Cos(float64(state.PlayerAngle)) * scale * speed
		dy := math.Sin(float64(state.PlayerAngle)) * scale * speed * -1
		m := g.State.Map

		if m.RetrieveTile(px+dx*3, py, g.Scale) > 0 {
			dx = 0
		}
		if m.RetrieveTile(px, py+dy*3, g.Scale) > 0 {
			dy = 0
		}

		state.PlayerPosition.X += dx
		state.PlayerPosition.Y += dy
	}
	if g.Inputs.S || g.Inputs.DOWN {
		speed := baseSpeed
		if g.Inputs.SHIFT {
			speed = runSpeed
		}
		px := state.PlayerPosition.X
		py := state.PlayerPosition.Y
		scale := float64(g.Scale)
		dx := math.Cos(float64(state.PlayerAngle)) * scale * speed * -1
		dy := math.Sin(float64(state.PlayerAngle)) * scale * speed
		m := g.State.Map

		if m.RetrieveTile(px+dx*3, py, g.Scale) > 0 {
			dx = 0
		}
		if m.RetrieveTile(px, py+dy*3, g.Scale) > 0 {
			dy = 0
		}

		state.PlayerPosition.X += dx
		state.PlayerPosition.Y += dy
	}
	if g.Inputs.A || g.Inputs.LEFT {
		if g.Inputs.ALT {
			speed := baseSpeed
			angle := state.PlayerAngle + (math.Pi / 2)
			if g.Inputs.SHIFT {
				speed = runSpeed
			}
			px := state.PlayerPosition.X
			py := state.PlayerPosition.Y
			scale := float64(g.Scale)
			dx := math.Cos(angle) * scale * speed * -1
			dy := math.Sin(angle) * scale * speed
			m := g.State.Map

			if m.RetrieveTile(px+dx*3, py, g.Scale) > 0 {
				dx = 0
			}
			if m.RetrieveTile(px, py+dy*3, g.Scale) > 0 {
				dy = 0
			}

			state.PlayerPosition.X += dx
			state.PlayerPosition.Y += dy
		} else {
			state.PlayerAngle -= baseAngleSpeed * (math.Pi / 180)
			if state.PlayerAngle < 0 {
				state.PlayerAngle += 2 * math.Pi
			}
		}
	}
	if g.Inputs.D || g.Inputs.RIGHT {
		if g.Inputs.ALT {
			speed := baseSpeed
			angle := state.PlayerAngle - (math.Pi / 2)
			if g.Inputs.SHIFT {
				speed = runSpeed
			}
			px := state.PlayerPosition.X
			py := state.PlayerPosition.Y
			scale := float64(g.Scale)
			dx := math.Cos(angle) * scale * speed * -1
			dy := math.Sin(angle) * scale * speed
			m := g.State.Map

			if m.RetrieveTile(px+dx*3, py, g.Scale) > 0 {
				dx = 0
			}
			if m.RetrieveTile(px, py+dy*3, g.Scale) > 0 {
				dy = 0
			}

			state.PlayerPosition.X += dx
			state.PlayerPosition.Y += dy
		} else {
			state.PlayerAngle += baseAngleSpeed * (math.Pi / 180)
			if state.PlayerAngle > (2 * math.Pi) {
				state.PlayerAngle -= 2 * math.Pi
			}
		}
	}
}
