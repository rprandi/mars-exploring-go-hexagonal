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

func Test_IsValidNewGrid_X_Zero_Invalid(t *testing.T) {
	xGrid := 0
	yGrid := 50
	grid, err := NewGrid(xGrid, yGrid)

	assert.Error(t, err, "Invalid grid")
	assert.Equal(t, grid, Grid{})
}

func Test_IsValidNewGrid_X_Negative_Invalid(t *testing.T) {
	xGrid := -1
	yGrid := 50
	grid, err := NewGrid(xGrid, yGrid)

	assert.Error(t, err, "Invalid grid")
	assert.Equal(t, grid, Grid{})
}

func Test_IsValidNewGrid_Y_Zero_Invalid(t *testing.T) {
	xGrid := 50
	yGrid := 0
	grid, err := NewGrid(xGrid, yGrid)

	assert.Error(t, err, "Invalid grid")
	assert.Equal(t, grid, Grid{})
}

func Test_IsValidNewGrid_Y_Negative_Invalid(t *testing.T) {
	xGrid := 50
	yGrid := -1
	grid, err := NewGrid(xGrid, yGrid)

	assert.Error(t, err, "Invalid grid")
	assert.Equal(t, grid, Grid{})
}
