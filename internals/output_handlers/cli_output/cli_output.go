package cli_output

import (
	"fmt"

	"github.com/rprandi/mars-exploring-go-hexagonal/internals/core/domain"
)

// Translate cli_input commands to the service

type OutputHandler struct {
}

func NewOutputHandler() OutputHandler {
	return OutputHandler{}
}

func (ch OutputHandler) WriteReport(mission domain.Mission) error {
	for _, probe := range mission.ReportMission() {
		fmt.Printf("%d %d %s\n", probe.CoordinateX, probe.CoordinateY, probe.Direction)
	}
	return nil
}
