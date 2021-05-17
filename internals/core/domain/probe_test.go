package domain

import (
	"testing"

	"gotest.tools/v3/assert"
)

const InvalidProbeInitialization = "invalid probe initialization"
const ProbeMovementInvalid = "probe movement is invalid"

func TestNewProbe_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "LRM"
	probe, err := NewProbe(xProbe, yProbe, North, commandsProbe)

	assert.NilError(t, err)
	assert.Equal(t, probe.CoordinateX, xProbe)
	assert.Equal(t, probe.CoordinateY, yProbe)
	assert.Equal(t, probe.Commands, commandsProbe)
	assert.Equal(t, probe.Direction, North)
}

func Test_Run_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "LRM"
	probe, err := NewProbe(xProbe, yProbe, North, commandsProbe)

	grid, _ := NewGrid(40, 40)

	probe.Run(grid)

	assert.NilError(t, err)
	assert.Equal(t, probe.CoordinateX, xProbe)
	assert.Equal(t, probe.CoordinateY, yProbe+1)
	assert.Equal(t, probe.Commands, commandsProbe)
	assert.Equal(t, probe.Direction, North)
}

func Test_Run_GridBoundaries_TopRight_Invalid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "LRM"
	probe, _ := NewProbe(xProbe, yProbe, North, commandsProbe)

	grid, _ := NewGrid(20, 30)

	err := probe.Run(grid)

	assert.Error(t, err, ProbeMovementInvalid)
}

func Test_Run_GridBoundaries_TopLeft_Invalid(t *testing.T) {
	xProbe := 0
	yProbe := 30
	commandsProbe := "LM"
	probe, _ := NewProbe(xProbe, yProbe, North, commandsProbe)

	grid, _ := NewGrid(20, 30)

	err := probe.Run(grid)

	assert.Error(t, err, ProbeMovementInvalid)
}

func Test_Run_GridBoundaries_BottomLeft_Invalid(t *testing.T) {
	xProbe := 0
	yProbe := 0
	commandsProbe := "LLM"
	probe, _ := NewProbe(xProbe, yProbe, North, commandsProbe)

	grid, _ := NewGrid(20, 30)

	err := probe.Run(grid)

	assert.Error(t, err, ProbeMovementInvalid)
}

func Test_Run_GridBoundaries_BottomRight_Invalid(t *testing.T) {
	xProbe := 20
	yProbe := 0
	commandsProbe := "RRM"
	probe, _ := NewProbe(xProbe, yProbe, North, commandsProbe)

	grid, _ := NewGrid(20, 30)

	err := probe.Run(grid)

	assert.Error(t, err, ProbeMovementInvalid)
}

func TestNewProbe_NoCommand_Invalid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := ""
	_, err := NewProbe(xProbe, yProbe, North, commandsProbe)

	assert.Error(t, err, InvalidProbeInitialization)
}

func TestNewProbe_CommandUnknown_Invalid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "X"
	_, err := NewProbe(xProbe, yProbe, North, commandsProbe)

	assert.Error(t, err, InvalidProbeInitialization)
}

func TestNewProbe_InvalidPosition_Invalid(t *testing.T) {
	xProbe := -1
	yProbe := 30
	commandsProbe := "LRM"
	_, err := NewProbe(xProbe, yProbe, North, commandsProbe)

	assert.Error(t, err, InvalidProbeInitialization)

	xProbe = 30
	yProbe = -1
	_, err = NewProbe(xProbe, yProbe, North, commandsProbe)

	assert.Error(t, err, InvalidProbeInitialization)

}

func TestNewProbe_InvalidDirection(t *testing.T) {
	xProbe := 30
	yProbe := 30
	commandsProbe := "LRM"

	_, err := NewProbe(xProbe, yProbe, "Q", commandsProbe)
	assert.Error(t, err, InvalidProbeInitialization)
}

func Test_IsValidCommand_Valid(t *testing.T) {
	commands := "LRM"
	valid := isValidCommands(commands)

	assert.Equal(t, valid, true)
}

func Test_IsValidCommand_WrongCommand_Invalid(t *testing.T) {
	commands := "LRXM"
	valid := isValidCommands(commands)

	assert.Equal(t, valid, false)
}

func Test_IsValidCommand_EmptyString_Invalid(t *testing.T) {
	commands := ""
	valid := isValidCommands(commands)

	assert.Equal(t, valid, false)
}

func Test_IsValidDirection_Valid(t *testing.T) {
	assert.Equal(t, isValidDirection(North), true)
	assert.Equal(t, isValidDirection(South), true)
	assert.Equal(t, isValidDirection(East), true)
	assert.Equal(t, isValidDirection(West), true)
}

func Test_IsValidDirection_Invalid(t *testing.T) {
	assert.Equal(t, isValidDirection("X"), false)
}

func Test_IsValidCoordinate_Valid(t *testing.T) {
	assert.Equal(t, isValidCoordinate(0), true)
	assert.Equal(t, isValidCoordinate(1), true)
	assert.Equal(t, isValidCoordinate(100), true)
}

func Test_IsValidCoordinate_Invalid(t *testing.T) {
	assert.Equal(t, isValidCoordinate(-1), false)
	assert.Equal(t, isValidCoordinate(-100), false)
}

func Test_MoveForward_North_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := North
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.moveForward()

	assert.NilError(t, err)
	assert.Equal(t, probe.CoordinateX, xProbe)
	assert.Equal(t, probe.CoordinateY, yProbe+1)
	assert.Equal(t, probe.Commands, commandsProbe)
	assert.Equal(t, probe.Direction, North)
}

func Test_MoveForward_South_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := South
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.moveForward()

	assert.NilError(t, err)
	assert.Equal(t, probe.CoordinateX, xProbe)
	assert.Equal(t, probe.CoordinateY, yProbe-1)
	assert.Equal(t, probe.Commands, commandsProbe)
	assert.Equal(t, probe.Direction, South)
}

func Test_MoveForward_East_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := East
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.moveForward()

	assert.NilError(t, err)
	assert.Equal(t, probe.CoordinateX, xProbe+1)
	assert.Equal(t, probe.CoordinateY, yProbe)
	assert.Equal(t, probe.Commands, commandsProbe)
	assert.Equal(t, probe.Direction, East)
}

func Test_MoveForward_West_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := West
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.moveForward()

	assert.NilError(t, err)
	assert.Equal(t, probe.CoordinateX, xProbe-1)
	assert.Equal(t, probe.CoordinateY, yProbe)
	assert.Equal(t, probe.Commands, commandsProbe)
	assert.Equal(t, probe.Direction, West)
}

func Test_MoveForward_InvalidDirection_Invalid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := "X"
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.moveForward()

	assert.Error(t, err, InvalidProbeInitialization)
}

func Test_NextStep_North_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := North
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	x, y := probe.NextStep()

	assert.NilError(t, err)
	assert.Equal(t, x, xProbe)
	assert.Equal(t, y, yProbe+1)
}

func Test_NextStep_South_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := South
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	x, y := probe.NextStep()

	assert.NilError(t, err)
	assert.Equal(t, x, xProbe)
	assert.Equal(t, y, yProbe-1)
}

func Test_NextStep_East_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := East
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	x, y := probe.NextStep()

	assert.NilError(t, err)
	assert.Equal(t, x, xProbe+1)
	assert.Equal(t, y, yProbe)
}

func Test_NextStep_West_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := West
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	x, y := probe.NextStep()

	assert.NilError(t, err)
	assert.Equal(t, x, xProbe-1)
	assert.Equal(t, y, yProbe)
}

func Test_RotateRight_FromNorth_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := North
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.rotateRight()

	assert.NilError(t, err)
	assert.Equal(t, probe.Direction, East)
}

func Test_RotateRight_FromSouth_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := South
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.rotateRight()

	assert.NilError(t, err)
	assert.Equal(t, probe.Direction, West)
}

func Test_RotateRight_FromEast_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := East
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.rotateRight()

	assert.NilError(t, err)
	assert.Equal(t, probe.Direction, South)
}

func Test_RotateRight_FromWest_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := West
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.rotateRight()

	assert.NilError(t, err)
	assert.Equal(t, probe.Direction, North)
}

func Test_RotateLeft_FromNorth_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := North
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.rotateLeft()

	assert.NilError(t, err)
	assert.Equal(t, probe.Direction, West)
}

func Test_RotateLeft_FromSouth_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := South
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.rotateLeft()

	assert.NilError(t, err)
	assert.Equal(t, probe.Direction, East)
}

func Test_RotateLeft_FromEast_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := East
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.rotateLeft()

	assert.NilError(t, err)
	assert.Equal(t, probe.Direction, North)
}

func Test_RotateLeft_FromWest_Valid(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "M"
	direction := West
	probe, err := NewProbe(xProbe, yProbe, direction, commandsProbe)

	probe.rotateLeft()

	assert.NilError(t, err)
	assert.Equal(t, probe.Direction, South)
}
