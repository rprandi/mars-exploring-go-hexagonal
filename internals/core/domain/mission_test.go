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

	err := mission.SetGridSize(30, 40)

	assert.NilError(t, err)
	assert.Equal(t, *mission.Grid, ValidGrid())
}

func TestMission_SetGrid_Invalid(t *testing.T) {
	mission := NewMission()

	err := mission.SetGridSize(-1, -1)

	assert.Error(t, err, "Invalid grid")
	assert.Equal(t, mission.State, OVER)
}

func TestMission_RunProbes_Valid(t *testing.T) {
	mission := NewMission()

	mission.SetGridSize(30, 40)
	mission.AddProbe(10, 10, North, "MMM")

	expectedProbe, _ := NewProbe(10, 13, North, "MMM")

	mission.RunProbes()

	assert.Equal(t, mission.State, READY)
	assert.Equal(t, *mission.Probes[0], expectedProbe)
}

func TestMission_RunProbes_AlreadyOver_Invalid(t *testing.T) {
	mission := NewMission()

	mission.State = OVER

	mission.SetGridSize(30, 40)
	mission.AddProbe(10, 10, North, "MMM")

	err := mission.RunProbes()

	assert.Error(t, err, "mission is over")
}

func TestMission_RunProbes_NoGrid_Invalid(t *testing.T) {
	mission := NewMission()

	mission.AddProbe(10, 10, North, "MMM")

	err := mission.RunProbes()

	assert.Error(t, err, "grid is missing")
}

func TestMission_RunProbes_NoProbes_Invalid(t *testing.T) {
	mission := NewMission()

	mission.State = OVER

	mission.SetGridSize(30, 40)

	err := mission.RunProbes()

	assert.Error(t, err, "mission is over")
}

func TestMission_RunProbes_ProbeCantRun_Invalid(t *testing.T) {
	mission := NewMission()

	mission.State = OVER

	mission.SetGridSize(30, 40)
	mission.AddProbe(30, 40, North, "MMM")

	err := mission.RunProbes()

	assert.Error(t, err, "mission is over")
}
