package domain

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewMission(t *testing.T) {
	xGrid := 40
	yGrid := 50
	grid := NewGrid(xGrid, yGrid)

	xProbe := 20
	yProbe := 30
	commandsProbe := "LRM"
	probe, err := NewProbe(xProbe, yProbe, North, commandsProbe)

	mission := NewMission(grid, []Probe{probe})

	assert.NilError(t, err)
	assert.Equal(t, mission.Probes[0], probe)
	assert.Equal(t, len(mission.Probes), 1)
	assert.Equal(t, mission.Plateau, grid)

}
