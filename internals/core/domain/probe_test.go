package domain

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewProbe_ValidProbe(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "LRM"
	probe, err := NewProbe(xProbe, yProbe, North, commandsProbe)

	assert.NilError(t, err)
	assert.Equal(t, probe.coordinateX, xProbe)
	assert.Equal(t, probe.coordinateY, yProbe)
	assert.Equal(t, probe.commands, commandsProbe)
}

func TestNewProbe_NoCommand(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := ""
	_, err := NewProbe(xProbe, yProbe, North, commandsProbe)

	assert.Error(t, err, "invalid command for Probe")
}

func TestNewProbe_CommandUnknown(t *testing.T) {
	xProbe := 20
	yProbe := 30
	commandsProbe := "X"
	_, err := NewProbe(xProbe, yProbe, North, commandsProbe)

	assert.Error(t, err, "invalid command for Probe")
}

func TestNewProbe_InvalidPosition(t *testing.T) {
	xProbe := -1
	yProbe := 30
	commandsProbe := "LRM"
	_, err := NewProbe(xProbe, yProbe, North, commandsProbe)

	assert.Error(t, err, "invalid coordinate for Probe")

	xProbe = 30
	yProbe = -1
	_, err = NewProbe(xProbe, yProbe, North, commandsProbe)

	assert.Error(t, err, "invalid coordinate for Probe")

}

func TestNewProbe_InvalidDirection(t *testing.T) {
	xProbe := 30
	yProbe := 30
	commandsProbe := "LRM"

	_, err := NewProbe(xProbe, yProbe, "Q", commandsProbe)
	assert.Error(t, err, "invalid direction for Probe")
}
