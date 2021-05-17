package filesystem_output

import (
	"fmt"
	"os"

	"github.com/rprandi/mars-exploring-go-hexagonal/internals/core/domain"
)

// Translate cli_input commands to the service

type OutputHandler struct {
	File string
}

func NewOutputHandler(file string) OutputHandler {
	return OutputHandler{File: file}
}

func (ch OutputHandler) WriteReport(mission domain.Mission) error {
	f, err := os.Create(ch.File)
	if err != nil {
		return err
	}

	defer f.Close()

	for _, probe := range mission.ReportMission() {
		reportLine := fmt.Sprintf("%d %d %s\n", probe.CoordinateX, probe.CoordinateY, probe.Direction)
		f.WriteString(reportLine)
	}
	return nil
}
