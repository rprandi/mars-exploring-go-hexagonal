package domain

import "errors"

// Structs das entities que podem ser usadas pela aplicação.
// Nem toda struct é um modelo de domínio, apenas as envolvidas na lógica de negócio são.

type Grid struct {
	horizontalSize int
	verticalSize   int
}

func NewGrid(horizontalSize int, verticalSize int) (Grid, error) {
	if IsValidInitialGridSize(horizontalSize, verticalSize) {
		return Grid{
			horizontalSize: horizontalSize,
			verticalSize:   verticalSize,
		}, nil
	}

	return Grid{}, errors.New("Invalid grid")
}

// returns true if within boundaries
func (g Grid) isValidProbeMovement(xProbe int, yProbe int) bool {
	//TODO is valid
	if xProbe <= g.horizontalSize && yProbe <= g.verticalSize && xProbe >= 0 && yProbe >= 0 {
		return true
	}

	return false
}

func IsValidInitialGridSize(x int, y int) bool {
	if (x > 1) && (y > 1) { //TODO teste inteiro
		return true
	}

	return false
}
