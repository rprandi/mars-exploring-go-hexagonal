package domain

// Structs das entities que podem ser usadas pela aplicação.
// Nem toda struct é um modelo de domínio, apenas as envolvidas na lógica de negócio são.

type Grid struct {
	horizontalSize int
	verticalSize   int
}

func NewGrid(horizontalSize int, verticalSize int) Grid {
	return Grid{
		horizontalSize: horizontalSize,
		verticalSize:   verticalSize,
	}
}
