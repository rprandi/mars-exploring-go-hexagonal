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
	coordinateX int
	coordinateY int
	direction   string
	commands    string
}

func NewProbe(coordinateX int, coordinateY int, direction string, commands string) (Probe, error) {
	err := isValidCoordinate(coordinateX)
	if err != nil {
		return Probe{}, err
	}

	err = isValidCoordinate(coordinateY)
	if err != nil {
		return Probe{}, err
	}

	err = isValidDirection(direction)
	if err != nil {
		return Probe{}, err
	}

	err = isValidCommands(commands)
	if err != nil {
		return Probe{}, err
	}

	return Probe{
		coordinateX: coordinateX,
		coordinateY: coordinateY,
		direction:   direction,
		commands:    commands,
	}, nil
}

func isValidCoordinate(coordinate int) error {
	if coordinate > 0 {
		return nil
	}

	return errors.New("invalid coordinate for Probe")
}

func isValidDirection(direction string) error {
	for _, validDirection := range ValidDirections {
		if validDirection == direction {
			return nil
		}
	}

	return errors.New("invalid direction for Probe")
}

// Delete all valid commands from the command string -> result should be empty
func isValidCommands(command string) error {
	if command == EmptyString {
		return errors.New("invalid command for Probe")
	}

	for _, ValidCommand := range ValidCommands {
		command = strings.ReplaceAll(command, ValidCommand, EmptyString)
	}

	if command == EmptyString { // Yes this comparison is duplicated, because it is post ReplaceAll
		return nil
	}

	return errors.New("invalid command for Probe")
}
