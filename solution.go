package square

import "math"

type typeSidesNumber int

const (
	SidesTriangle typeSidesNumber = 3
	SidesSquare   typeSidesNumber = 4
	SidesCircle   typeSidesNumber = 0
)

func CalcSquare(sideLen float64, sidesNum typeSidesNumber) float64 {
	totalSquare := math.Pow(sideLen, 2)
	switch sidesNum {
	case SidesTriangle:
		return totalSquare * math.Sqrt(3) / 4
	case SidesSquare:
		return totalSquare
	case SidesCircle:
		return totalSquare * math.Pi
	}
	return 0
}
