package ports

// Interfaces pra comunicar com os actors

type MissionService interface {
	CreateMission()
	SetGrid(int, int) error
	AddProbe(int, int, string, string) error
	RunProbes() error
	ReportMission() error
}
