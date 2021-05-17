package domain

import (
	"testing"

	"gotest.tools/v3/assert"
)

func ValidGrid() Grid {
	grid, _ := NewGrid(30, 40)
	return grid
}

func ValidProbe() Probe {
	xProbe := 20
	yProbe := 30
	commandsProbe := "LRM"
	probe, _ := NewProbe(xProbe, yProbe, North, commandsProbe)
	return probe
}

func TestNewMission(t *testing.T) {
	mission := NewMission()

	assert.Equal(t, mission.State, READY)
}

func TestMission_SetGrid_Valid(t *testing.T) {
	mission := NewMission()

	grid := ValidGrid()

	err := mission.SetGrid(grid)

	assert.NilError(t, err)
	assert.Equal(t, *mission.Grid, ValidGrid())
}

func TestMission_RunProbes_Valid(t *testing.T) {
	mission := NewMission()

	grid := ValidGrid()
	probe := ValidProbe()

	_ = mission.SetGrid(grid)
	mission.AddProbe(probe)

	expectedProbe, _ := NewProbe(20, 31, "N", "LRM")

	mission.RunProbes()

	assert.Equal(t, mission.State, READY)
	assert.Equal(t, *mission.Probes[0], expectedProbe)
}

func TestMission_RunProbes_AlreadyOver_Invalid(t *testing.T) {
	mission := NewMission()

	mission.State = OVER

	grid := ValidGrid()
	probe := ValidProbe()

	_ = mission.SetGrid(grid)
	mission.AddProbe(probe)

	err := mission.RunProbes()

	assert.Error(t, err, "mission is over")
}

func TestMission_RunProbes_NoGrid_Invalid(t *testing.T) {
	mission := NewMission()

	probe := ValidProbe()

	mission.AddProbe(probe)

	err := mission.RunProbes()

	assert.Error(t, err, "grid is missing")
}

func TestMission_RunProbes_NoProbes_Invalid(t *testing.T) {
	mission := NewMission()

	mission.State = OVER

	grid := ValidGrid()

	_ = mission.SetGrid(grid)

	err := mission.RunProbes()

	assert.Error(t, err, "mission is over")
}

func TestMission_RunProbes_ProbeCantRun_Invalid(t *testing.T) {
	mission := NewMission()

	mission.State = OVER

	grid := ValidGrid()
	probe := ValidProbe()

	_ = mission.SetGrid(grid)
	mission.AddProbe(probe)

	err := mission.RunProbes()

	assert.Error(t, err, "mission is over")
}
