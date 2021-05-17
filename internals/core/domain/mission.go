package domain

import "errors"

// Structs das entities que podem ser usadas pela aplicação.
// Nem toda struct é um modelo de domínio, apenas as envolvidas na lógica de negócio são.

// A Mission has a plateau and a set of probes
const READY = "READY"
const OVER = "OVER"

type Mission struct {
	Grid   *Grid
	Probes []*Probe
	State  string
}

// TODO collision of probes
func NewMission() Mission {
	return Mission{State: READY}
}

func (m *Mission) SetGrid(grid Grid) error {
	m.Grid = &grid

	return nil
}

func (m *Mission) AddProbe(probe Probe) error {
	m.Probes = append(m.Probes, &probe)

	return nil
}

func (m *Mission) RunProbes() error {
	if m.IsOver() {
		return errors.New("mission is over")
	}

	if m.Grid == nil {
		return errors.New("grid is missing")
	}

	for _, probe := range m.Probes {
		err := probe.Run(*m.Grid)
		if err != nil {
			m.State = "OVER"
			return err
		}
	}

	return nil
}

func (m Mission) IsOver() bool {
	if m.State == OVER {
		return true
	}
	return false
}

func (m *Mission) ReportMission() []*Probe {
	return m.Probes
}
