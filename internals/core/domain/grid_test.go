package domain

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewGrid(t *testing.T) {
	xGrid := 40
	yGrid := 50
	grid := NewGrid(xGrid, yGrid)

	assert.Equal(t, grid.horizontalSize, xGrid)
	assert.Equal(t, grid.verticalSize, yGrid)
}
