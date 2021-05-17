package main

import (
	"github.com/rprandi/mars-exploring-go-hexagonal/internals/core/services/mission_service"
	"github.com/rprandi/mars-exploring-go-hexagonal/internals/input_handlers/filesystem_input"
	"github.com/rprandi/mars-exploring-go-hexagonal/internals/output_handlers/filesystem_output"
)

func main() {

	// Both handlers below handle terminal/cli input/output.  Uncomment the 3 lines below for Terminal/CLI
	//CLIInputHandler := cli_input.NewInputCLIHandler()
	//CLIOutputHandler := cli_output.NewOutputHandler()
	//missionService := mission_service.NewMissionService(CLIInputHandler, CLIOutputHandler)

	// Both handlers below handle file input/output.  Uncomment the 3 lines below for file input/output
	FileInputHandler := filesystem_input.NewInputHandler("input.txt")
	FileOutputHandler := filesystem_output.NewOutputHandler("output.txt")
	missionService := mission_service.NewMissionService(FileInputHandler, FileOutputHandler)

	missionService.CreateMission()
	missionService.RunMission()
	missionService.ReportMission()
}
