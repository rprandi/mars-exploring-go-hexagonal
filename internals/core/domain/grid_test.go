package domain

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_NewGrid_Valid(t *testing.T) {
	xGrid := 40
	yGrid := 50
	grid, err := NewGrid(xGrid, yGrid)

	assert.NilError(t, err)
	assert.Equal(t, grid.horizontalSize, xGrid)
	assert.Equal(t, grid.verticalSize, yGrid)
}

func Test_NewGrid_X_Zero_Invalid(t *testing.T) {
	xGrid := 0
	yGrid := 50
	grid, err := NewGrid(xGrid, yGrid)

	assert.Error(t, err, "Invalid grid")
	assert.Equal(t, grid, Grid{})
}

func Test_NewGrid_X_Negative_Invalid(t *testing.T) {
	xGrid := -1
	yGrid := 50
	grid, err := NewGrid(xGrid, yGrid)

	assert.Error(t, err, "Invalid grid")
	assert.Equal(t, grid, Grid{})
}

func Test_NewGrid_Y_Zero_Invalid(t *testing.T) {
	xGrid := 50
	yGrid := 0
	grid, err := NewGrid(xGrid, yGrid)

	assert.Error(t, err, "Invalid grid")
	assert.Equal(t, grid, Grid{})
}

func Test_NewGrid_Y_Negative_Invalid(t *testing.T) {
	xGrid := 50
	yGrid := -1
	grid, err := NewGrid(xGrid, yGrid)

	assert.Error(t, err, "Invalid grid")
	assert.Equal(t, grid, Grid{})
}

func Test_IsValidProbeMovement_Valid(t *testing.T) {
	grid, _ := NewGrid(50, 50)

	assert.Equal(t, grid.isValidProbeMovement(0, 0), true)
	assert.Equal(t, grid.isValidProbeMovement(0, 50), true)
	assert.Equal(t, grid.isValidProbeMovement(50, 50), true)
	assert.Equal(t, grid.isValidProbeMovement(50, 0), true)
}

func Test_IsValidProbeMovement_Invalid(t *testing.T) {
	grid, _ := NewGrid(50, 50)

	assert.Equal(t, grid.isValidProbeMovement(-1, 0), false)
	assert.Equal(t, grid.isValidProbeMovement(0, -1), false)
	assert.Equal(t, grid.isValidProbeMovement(0, 51), false)
	assert.Equal(t, grid.isValidProbeMovement(51, 51), false)
	assert.Equal(t, grid.isValidProbeMovement(51, 0), false)
}

func Test_IsValidInitialGridSize_Valid(t *testing.T) {
	assert.Equal(t, IsValidInitialGridSize(1, 1), true)
	assert.Equal(t, IsValidInitialGridSize(10, 1), true)
	assert.Equal(t, IsValidInitialGridSize(1, 10), true)
	assert.Equal(t, IsValidInitialGridSize(10, 10), true)
}

func Test_IsValidInitialGridSize_Invalid(t *testing.T) {
	assert.Equal(t, IsValidInitialGridSize(0, 1), false)
	assert.Equal(t, IsValidInitialGridSize(1, 0), false)
	assert.Equal(t, IsValidInitialGridSize(-1, 10), false)
	assert.Equal(t, IsValidInitialGridSize(10, -1), false)
}
