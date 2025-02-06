package engine

type MoveCoordinate struct {
	X int8
	Y int8
}

func (coord MoveCoordinate) isValid() bool {
	return coord.X >= 0 && coord.X < 4 && coord.Y >= 0 && coord.Y < 4
}
