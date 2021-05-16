package domain

// Structs das entities que podem ser usadas pela aplicação.
// Nem toda struct é um modelo de domínio, apenas as envolvidas na lógica de negócio são.

const (
	NORTH        = 'N'
	SOUTH        = 'S'
	EAST         = 'E'
	WEST         = 'W'
	TURN_RIGHT   = 'R'
	TURN_LEFT    = 'L'
	MOVE_FORWARD = 'M'
)

type Probe struct {
	coordinateX int
	coordinateY int
	direction   string
	commands    string
}

func NewProbe(coordinateX int, coordinateY int, direction string, commands string) Probe {
	return Probe{
		coordinateX: coordinateX,
		coordinateY: coordinateY,
		direction:   direction,
		commands:    commands,
	}
}
