package cli_input

import (
	"fmt"

	"github.com/rprandi/mars-exploring-go-hexagonal/internals/core/domain"

	"github.com/rprandi/mars-exploring-go-hexagonal/internals/core/ports"
)

// Translate cli_input commands to the service

type InputHandler struct {
	MissionService ports.MissionService
}

func NewInputCLIHandler() *InputHandler {
	return &InputHandler{}
}

func (ch *InputHandler) ReadGrid() (domain.Grid, error) {
	var grid domain.Grid
	var xGrid int
	var yGrid int

	_, err := fmt.Scanf("%d", &xGrid)
	if err != nil {
		return domain.Grid{}, nil
	}

	_, err = fmt.Scanf("%d", &yGrid)
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
	var probes []domain.Probe

	for {
		var xProbe int
		var yProbe int
		var direction string
		var command string

		_, err := fmt.Scanf("%d", &xProbe)
		if err != nil {
			break
		}

		_, err = fmt.Scanf("%d", &yProbe)
		if err != nil {
			break
		}

		_, err = fmt.Scanf("%s", &direction)
		if err != nil {
			break
		}

		_, err = fmt.Scanf("%s", &command)
		if err != nil {
			break
		}

		probe, _ := domain.NewProbe(xProbe, yProbe, direction, command)
		probes = append(probes, probe)
	}

	return probes, nil
}
