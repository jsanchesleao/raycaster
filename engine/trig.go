package engine

import "math"

func Tan(angle float64) float64 {
	var tan float64
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	if cos > 0 && cos < 0.001 {
		tan = sin / 0.001
	} else if cos < 0 && cos > -0.001 {
		tan = sin / -0.001
	} else if cos == 0 {
		tan = 10
	} else {
		tan = sin / cos
	}
	return tan
}

func ATan(angle float64) float64 {
	var aTan float64
	cos := math.Cos(angle)
	sin := math.Sin(angle)

	if sin > 0 && sin < 0.001 {
		aTan = cos / 0.001
	} else if sin < 0 && sin > -0.001 {
		aTan = cos / -0.001
	} else if sin == 0 {
		aTan = 10
	} else {
		aTan = cos / sin
	}

	return aTan
}
