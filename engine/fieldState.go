package engine

type FieldState byte

const (
	Empty FieldState = iota
	White
	Black
)

func (state FieldState) GetString() string {
	switch state {
	case Empty:
		return "_"
	case White:
		return "O"
	case Black:
		return "X"
	}
	return "?"
}

func (state FieldState) GetName() string {
	switch state {
	case Empty:
		return "Empty"
	case White:
		return "White"
	case Black:
		return "Black"
	}
	return "???"
}

func (state FieldState) Flip() FieldState {
	switch state {
	case White:
		return Black
	case Black:
		return White
	default:
		return state
	}
}
