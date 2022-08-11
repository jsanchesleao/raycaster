package engine

import "math"

func MovePlayer(g *Game) {
	state := &g.State
	if g.Inputs.W || g.Inputs.UP {
		dx := 2 * math.Cos(float64(state.PlayerAngle))
		dy := 2 * math.Sin(float64(state.PlayerAngle))
		nextTile := g.State.Map.RetrieveTile(
			state.PlayerPosition.X+4*dx,
			state.PlayerPosition.Y-4*dy,
			g.Scale,
		)
		if nextTile == 0 {
			state.PlayerPosition.X += dx
			state.PlayerPosition.Y -= dy
		}
	}
	if g.Inputs.S || g.Inputs.DOWN {
		dx := 2 * math.Cos(float64(state.PlayerAngle))
		dy := 2 * math.Sin(float64(state.PlayerAngle))
		nextTile := state.Map.RetrieveTile(
			state.PlayerPosition.X-4*dx,
			state.PlayerPosition.Y+4*dy,
			g.Scale,
		)
		if nextTile == 0 {
			state.PlayerPosition.X -= dx
			state.PlayerPosition.Y += dy
		}
	}
	if g.Inputs.D || g.Inputs.RIGHT {
		state.PlayerAngle -= 2 * (math.Pi / 180)
		if state.PlayerAngle < 0 {
			state.PlayerAngle += 2 * math.Pi
		}
	}
	if g.Inputs.A || g.Inputs.LEFT {
		state.PlayerAngle += 2 * (math.Pi / 180)
		if state.PlayerAngle > (2 * math.Pi) {
			state.PlayerAngle -= 2 * math.Pi
		}
	}
}
