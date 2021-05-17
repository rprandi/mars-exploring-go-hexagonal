package ports

import "github.com/rprandi/mars-exploring-go-hexagonal/internals/core/domain"

// Interfaces pra comunicar com os actors

type MissionService interface {
	CreateMission() error
	RunMission() error
	ReportMission()
}

type InputHandler interface {
	ReadGrid() (domain.Grid, error)
	ReadProbes() ([]domain.Probe, error)
}

type OutputHandler interface {
	WriteReport(domain.Mission) error
}
