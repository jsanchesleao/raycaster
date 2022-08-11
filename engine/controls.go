package engine

import "math"

const baseSpeed = 0.02
const baseAngleSpeed = 2

func MovePlayer(g *Game) {
	state := &g.State
	if g.Inputs.W || g.Inputs.UP {
		dx := 2 * math.Cos(float64(state.PlayerAngle))
		dy := 2 * math.Sin(float64(state.PlayerAngle))
		nextTile := g.State.Map.RetrieveTile(
			state.PlayerPosition.X+float64(g.Scale)*baseSpeed*8*dx,
			state.PlayerPosition.Y-float64(g.Scale)*baseSpeed*8*dy,
			g.Scale,
		)
		if nextTile == 0 {
			state.PlayerPosition.X += float64(g.Scale) * baseSpeed * dx
			state.PlayerPosition.Y -= float64(g.Scale) * baseSpeed * dy
		}
	}
	if g.Inputs.S || g.Inputs.DOWN {
		dx := 2 * math.Cos(float64(state.PlayerAngle))
		dy := 2 * math.Sin(float64(state.PlayerAngle))
		nextTile := state.Map.RetrieveTile(
			state.PlayerPosition.X-float64(g.Scale)*baseSpeed*8*dx,
			state.PlayerPosition.Y+float64(g.Scale)*baseSpeed*8*dy,
			g.Scale,
		)
		if nextTile == 0 {
			state.PlayerPosition.X -= float64(g.Scale) * baseSpeed * dx
			state.PlayerPosition.Y += float64(g.Scale) * baseSpeed * dy
		}
	}
	if g.Inputs.A || g.Inputs.LEFT {
		state.PlayerAngle -= baseAngleSpeed * (math.Pi / 180)
		if state.PlayerAngle < 0 {
			state.PlayerAngle += 2 * math.Pi
		}
	}
	if g.Inputs.D || g.Inputs.RIGHT {
		state.PlayerAngle += baseAngleSpeed * (math.Pi / 180)
		if state.PlayerAngle > (2 * math.Pi) {
			state.PlayerAngle -= 2 * math.Pi
		}
	}
}
