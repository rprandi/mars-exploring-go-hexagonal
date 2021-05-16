package domain

import (
	"errors"
	"strings"
)

// Structs das entities que podem ser usadas pela aplicação.
// Nem toda struct é um modelo de domínio, apenas as envolvidas na lógica de negócio são.

const (
	North       = "N"
	South       = "S"
	East        = "E"
	West        = "W"
	TurnRight   = "R"
	TurnLeft    = "L"
	MoveForward = "M"
	EmptyString = ""
)

var ValidDirections = []string{North, South, East, West}
var ValidCommands = []string{TurnRight, TurnLeft, MoveForward}

type Probe struct {
	CoordinateX int
	CoordinateY int
	Direction   string
	Commands    string
}

func NewProbe(coordinateX int, coordinateY int, direction string, commands string) (Probe, error) {
	if isValidCoordinate(coordinateX) && isValidCoordinate(coordinateY) && isValidDirection(direction) && isValidCommands(commands) {
		return Probe{
			CoordinateX: coordinateX,
			CoordinateY: coordinateY,
			Direction:   direction,
			Commands:    commands,
		}, nil
	}

	return Probe{}, errors.New("invalid probe initialization")
}

// For each command, check if it is valid and execute
func (p *Probe) Run(grid Grid) error {
	for _, command := range strings.Split(p.Commands, EmptyString) {
		switch command {
		case TurnRight:
			p.rotateRight()
		case TurnLeft:
			p.rotateLeft()
		case MoveForward:
			x, y := p.NextStep()
			if grid.isValidProbeMovement(x, y) {
				p.moveForward()
				continue
			}
			return errors.New("probe movement is invalid")
		}
	}
	return nil
}

func (p *Probe) rotateRight() {
	switch p.Direction {
	case North:
		p.Direction = East
	case South:
		p.Direction = West
	case East:
		p.Direction = South
	case West:
		p.Direction = North
	}
}

func (p *Probe) rotateLeft() {
	switch p.Direction {
	case North:
		p.Direction = West
	case South:
		p.Direction = East
	case East:
		p.Direction = North
	case West:
		p.Direction = South
	}
}

func (p *Probe) NextStep() (int, int) {
	switch p.Direction {
	case North:
		return p.CoordinateX, p.CoordinateY + 1
	case South:
		return p.CoordinateX, p.CoordinateY - 1
	case East:
		return p.CoordinateX + 1, p.CoordinateY
	case West:
		return p.CoordinateX - 1, p.CoordinateY
	}

	return -1, -1
}

func (p *Probe) moveForward() {
	switch p.Direction {
	case North:
		p.CoordinateY += 1
	case South:
		p.CoordinateY -= 1
	case East:
		p.CoordinateX += 1
	case West:
		p.CoordinateX -= 1
	}
}

func isValidCoordinate(coordinate int) bool {
	if coordinate >= 0 {
		return true
	}

	return false
}

func isValidDirection(direction string) bool {
	for _, validDirection := range ValidDirections {
		if validDirection == direction {
			return true
		}
	}

	return false
}

// Delete all valid commands from the command string -> result should be empty
func isValidCommands(command string) bool {
	if command == EmptyString {
		return false
	}

	for _, ValidCommand := range ValidCommands {
		command = strings.ReplaceAll(command, ValidCommand, EmptyString)
	}

	if command == EmptyString { // Yes this comparison is duplicated, because it is post ReplaceAll
		return true
	}

	return false
}
