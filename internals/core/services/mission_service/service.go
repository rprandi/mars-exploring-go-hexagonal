package mission_service

import "github.com/rprandi/mars-exploring-go-hexagonal/internals/core/domain"

// Entrada para o core, cada um implementa a porta

type MissionService struct {
	Mission domain.Mission
}

func NewMissionService() MissionService {
	return MissionService{}
}

func (ms *MissionService) CreateMission() {
	mission := domain.NewMission()

	ms.Mission = mission
}

func (ms *MissionService) SetGrid(x int, y int) error {
	err := ms.Mission.SetGridSize(x, y)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MissionService) AddProbe(x int, y int, direction string, commands string) error {
	err := ms.Mission.AddProbe(x, y, direction, commands)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MissionService) RunProbes() error {
	err := ms.Mission.RunProbes()
	if err != nil {
		return err
	}
	return nil
}

func (ms *MissionService) ReportMission() error {
	// TODO How to report the mission ?
	return nil
}
