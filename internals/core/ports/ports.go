package ports

// Interfaces pra comunicar com os actors

type MissionService interface {
	CreateMission(int, int) error
	AddProbe(int, int, string, string) error
	RunProbes() error
}

//
//
//
//- Create a new grid
//
//Given a coordinate (X, Y) that is the top-right corner it should create a new grid. It should validate the size of the grid to only allow positive integers.
//
//- Place probe
//
//Given a coordinate (X, Y) it should accept a probe in that position. It should check for grid boundaries and other probes' collision.
//
//- Turn probe
//
//Given a probe and an direction (L, R), it should rotate the probe 90 degrees to a new direction
//
//- Move probe
//
//Given a probe, it should move it one step in the direction it is facing. It should check for grid boundaries and other probes' collision.
//
//- Get probe
//
//Given a probe, it should return the current position and direction.
