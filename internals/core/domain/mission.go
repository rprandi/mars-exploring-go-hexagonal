package domain

// Structs das entities que podem ser usadas pela aplicação.
// Nem toda struct é um modelo de domínio, apenas as envolvidas na lógica de negócio são.

// A Mission has a plateau and a set of probes
type Mission struct {
	Plateau Grid
	Probes  []Probe
}

func NewMission(grid Grid, probes []Probe) Mission {
	return Mission{
		Plateau: grid,
		Probes:  probes,
	}
}
