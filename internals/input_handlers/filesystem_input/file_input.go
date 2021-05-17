package filesystem_input

import (
	"fmt"
	"os"

	"github.com/rprandi/mars-exploring-go-hexagonal/internals/core/domain"

	"github.com/rprandi/mars-exploring-go-hexagonal/internals/core/ports"
)

// Translate cli_input commands to the service

type InputHandler struct {
	MissionService ports.MissionService
	File           string
}

func NewInputCLIHandler(file string) *InputHandler {
	return &InputHandler{File: file}
}

func (ch *InputHandler) ReadGrid() (domain.Grid, error) {
	f, err := os.Open(ch.File)
	if err != nil {
		return domain.Grid{}, nil
	}
	defer f.Close()

	var grid domain.Grid
	var xGrid int
	var yGrid int

	_, err = fmt.Fscanf(f, "%d", &xGrid)
	if err != nil {
		return domain.Grid{}, nil
	}

	_, err = fmt.Fscanf(f, "%d", &yGrid)
	if err != nil {
		return domain.Grid{}, nil
	}

	grid, err = domain.NewGrid(xGrid, yGrid)
	if err != nil {
		return domain.Grid{}, nil
	}

	return grid, nil
}

func (ch *InputHandler) ReadProbes() ([]domain.Probe, error) {
	f, err := os.Open(ch.File)
	if err != nil {
		return []domain.Probe{}, nil
	}
	defer f.Close()

	var probes []domain.Probe

	for {
		var xProbe int
		var yProbe int
		var direction string
		var command string

		_, err := fmt.Fscanf(f, "%d", &xProbe)
		if err != nil {
			break
		}

		_, err = fmt.Fscanf(f, "%d", &yProbe)
		if err != nil {
			break
		}

		_, err = fmt.Fscanf(f, "%s", &direction)
		if err != nil {
			break
		}

		_, err = fmt.Fscanf(f, "%s", &command)
		if err != nil {
			break
		}

		probe, _ := domain.NewProbe(xProbe, yProbe, direction, command)
		probes = append(probes, probe)
	}

	return probes, nil
}
