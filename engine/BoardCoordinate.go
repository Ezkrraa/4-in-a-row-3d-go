package engine

type BoardCoordinate struct {
	X int8
	Y int8
	Z int8
}

func (coord BoardCoordinate) isValid() bool {
	return coord.X >= 0 && coord.X < 4 && coord.Y >= 0 && coord.Y < 4 && coord.Z >= 0 && coord.Z < 4
}
