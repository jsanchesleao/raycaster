package engine

import (
	"math"
)

type Tile int

type Map struct {
	Width int
	Tiles []Tile
	Start Position
}

type Position struct {
	X, Y float64
}

type GameState struct {
	Map            *Map
	PlayerPosition *Position
	PlayerAngle    Angle
}

type Ray struct {
	Origin Position
	Angle  Angle
	Length float64
	Detect Tile // tile it detected
}

type Angle = float64

func Radians(deg int) Angle {
	if deg < 0 {
		deg = 360 - (-deg)%360
	} else if deg > 360 {
		deg = deg % 360
	}

	return (float64(deg) * math.Pi) / 180
}

func (m *Map) RetrieveTile(x, y float64, scale int) Tile {
	row := (int(y) / scale)
	col := (int(x) / scale)

	index := row*m.Width + col
	return m.Tiles[index]
}
