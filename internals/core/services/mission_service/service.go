package mission_service

import (
	"github.com/rprandi/mars-exploring-go-hexagonal/internals/core/ports"

	"github.com/rprandi/mars-exploring-go-hexagonal/internals/core/domain"
)

// Entrada para o core

type MissionService struct {
	Mission       domain.Mission
	InputHandler  ports.InputHandler
	OutputHandler ports.OutputHandler
}

func NewMissionService(inputHandler ports.InputHandler, outputHandler ports.OutputHandler) *MissionService {
	return &MissionService{
		InputHandler:  inputHandler,
		OutputHandler: outputHandler,
	}
}

func (ms *MissionService) CreateMission() error {
	mission := domain.NewMission()

	ms.Mission = mission

	grid, err := ms.InputHandler.ReadGrid()
	if err != nil {
		return err
	}

	ms.Mission.SetGrid(grid)

	probes, err := ms.InputHandler.ReadProbes()
	if err != nil {
		return err
	}

	for _, probe := range probes {
		ms.Mission.AddProbe(probe)
	}

	return nil
}

func (ms *MissionService) SetGrid(grid domain.Grid) error {
	err := ms.Mission.SetGrid(grid)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MissionService) AddProbe(probe domain.Probe) error {
	err := ms.Mission.AddProbe(probe)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MissionService) RunMission() error {
	err := ms.Mission.RunProbes()
	if err != nil {
		return err
	}
	return nil
}

func (ms *MissionService) ReportMission() {
	ms.OutputHandler.WriteReport(ms.Mission)
}
