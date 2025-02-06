package engine

type GameState struct {
	Board         [4][4][4]FieldState
	CurrentPlayer FieldState
	MoveHistory   []MoveCoordinate
}

func CreateEmpty() GameState {
	return GameState{
		[4][4][4]FieldState{
			{
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
			}, {
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
			}, {
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
			}, {
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
				{Empty, Empty, Empty, Empty},
			},
		},
		White,
		[]MoveCoordinate{},
	}
}

func (state *GameState) GetMovedClone(coordinate MoveCoordinate) *GameState {
	newState := GameState{state.Board, state.CurrentPlayer, state.MoveHistory}
	Z := state.findSpot(coordinate)
	if Z < 0 {
		panic("Invalid spot for move")
	}
	newState.Board[Z][coordinate.Y][coordinate.X] = newState.CurrentPlayer
	newState.CurrentPlayer = newState.CurrentPlayer.Flip()
	newState.MoveHistory = append(newState.MoveHistory, coordinate)
	return &newState
}

// returns whether a winner was found, and who it was
func (state *GameState) GetWinner() (bool, FieldState) {
	// check lines
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			if isLineWon(state.Board[0][a][b], state.Board[1][a][b], state.Board[2][a][b], state.Board[3][a][b]) {
				return true, state.Board[0][a][b]
			} else if isLineWon(state.Board[a][0][b], state.Board[a][1][b], state.Board[a][2][b], state.Board[a][3][b]) {
				return true, state.Board[a][0][b]
			} else if isLineWon(state.Board[a][b][0], state.Board[a][b][1], state.Board[a][b][2], state.Board[a][b][3]) {
				return true, state.Board[a][b][0]
			}
		}
	}

	// 2d diagonals
	for d := 0; d < 4; d++ {
		if isLineWon(state.Board[d][0][0], state.Board[d][1][1], state.Board[d][2][2], state.Board[d][3][3]) {
			return true, state.Board[d][0][0]
		} else if isLineWon(state.Board[d][0][3], state.Board[d][1][2], state.Board[d][2][1], state.Board[d][3][0]) {
			return true, state.Board[d][0][3]
		} else if isLineWon(state.Board[0][d][0], state.Board[1][d][1], state.Board[2][d][2], state.Board[3][d][3]) {
			return true, state.Board[0][d][0]
		} else if isLineWon(state.Board[0][d][3], state.Board[1][d][2], state.Board[2][d][1], state.Board[3][d][0]) {
			return true, state.Board[0][d][3]
		} else if isLineWon(state.Board[0][0][d], state.Board[1][1][d], state.Board[2][2][d], state.Board[3][3][d]) {
			return true, state.Board[0][0][d]
		} else if isLineWon(state.Board[0][3][d], state.Board[1][2][d], state.Board[2][1][d], state.Board[3][0][d]) {
			return true, state.Board[0][3][d]
		}
	}

	// 3d diagonals
	if isLineWon(state.Board[0][0][0], state.Board[1][1][1], state.Board[2][2][2], state.Board[3][3][3]) {
		return true, state.Board[0][0][0]
	} else if isLineWon(state.Board[0][0][3], state.Board[1][1][2], state.Board[2][2][1], state.Board[3][3][0]) {
		return true, state.Board[0][0][3]
	} else if isLineWon(state.Board[0][3][0], state.Board[1][2][1], state.Board[2][1][2], state.Board[3][0][3]) {
		return true, state.Board[0][3][0]
	} else if isLineWon(state.Board[0][3][3], state.Board[1][2][2], state.Board[2][1][1], state.Board[3][0][0]) {
		return true, state.Board[0][3][3]
	}
	// nothing
	for z := 0; z < 4; z++ {
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				if state.Board[z][y][x] == Empty {
					return false, Empty
				}
			}
		}
	}
	return true, Empty
}

func isLineWon(n0 FieldState, n1 FieldState, n2 FieldState, n3 FieldState) bool {
	return n0 != Empty && n0 == n1 && n1 == n2 && n2 == n3
}

func (state *GameState) makeMove(coordinate MoveCoordinate) bool {
	spot := state.findSpot(coordinate)
	if spot == -1 {
		return false
	}
	state.Board[spot][coordinate.Y][coordinate.X] = state.CurrentPlayer
	state.CurrentPlayer = state.CurrentPlayer.Flip()
	return true
}

// sees if there is an empty spot at that coordinate, returns the z index
func (state *GameState) findSpot(coordinate MoveCoordinate) int8 {
	for z := int8(0); z < 4; z++ {
		if state.Board[z][coordinate.Y][coordinate.X] == Empty {
			return z
		}
	}
	return -1
}

func (state GameState) IsValid() bool {
	whiteFound, blackFound := 0, 0
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			supported := true
			for z := 0; z < 4; z++ {
				if state.Board[z][y][x] != Empty {
					if !supported {
						return false
					} else if state.Board[z][y][x] == White {
						whiteFound++
					} else {
						blackFound++
					}
				}
			}
		}
	}
	diff := whiteFound - blackFound
	return diff == 0 || diff == 1
}

// returns whether a winner was found, and who it was
func (state *GameState) CountNearWins(player FieldState) float32 {
	count := float32(0)
	// check lines
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			count += countEqualTo(player, []FieldState{state.Board[0][a][b], state.Board[1][a][b], state.Board[2][a][b], state.Board[3][a][b]})
			count += countEqualTo(player, []FieldState{state.Board[a][0][b], state.Board[a][1][b], state.Board[a][2][b], state.Board[a][3][b]})
			count += countEqualTo(player, []FieldState{state.Board[a][b][0], state.Board[a][b][1], state.Board[a][b][2], state.Board[a][b][3]})
		}
	}

	// 2d diagonals
	for d := 0; d < 4; d++ {
		count += countEqualTo(player, []FieldState{state.Board[d][0][0], state.Board[d][1][1], state.Board[d][2][2], state.Board[d][3][3]})
		count += countEqualTo(player, []FieldState{state.Board[d][0][3], state.Board[d][1][2], state.Board[d][2][1], state.Board[d][3][0]})
		count += countEqualTo(player, []FieldState{state.Board[0][d][0], state.Board[1][d][1], state.Board[2][d][2], state.Board[3][d][3]})
		count += countEqualTo(player, []FieldState{state.Board[0][d][3], state.Board[1][d][2], state.Board[2][d][1], state.Board[3][d][0]})
		count += countEqualTo(player, []FieldState{state.Board[0][0][d], state.Board[1][1][d], state.Board[2][2][d], state.Board[3][3][d]})
		count += countEqualTo(player, []FieldState{state.Board[0][3][d], state.Board[1][2][d], state.Board[2][1][d], state.Board[3][0][d]})
	}

	// 3d diagonals
	count += countEqualTo(player, []FieldState{state.Board[0][0][0], state.Board[1][1][1], state.Board[2][2][2], state.Board[3][3][3]})
	count += countEqualTo(player, []FieldState{state.Board[0][0][3], state.Board[1][1][2], state.Board[2][2][1], state.Board[3][3][0]})
	count += countEqualTo(player, []FieldState{state.Board[0][3][0], state.Board[1][2][1], state.Board[2][1][2], state.Board[3][0][3]})
	count += countEqualTo(player, []FieldState{state.Board[0][3][3], state.Board[1][2][2], state.Board[2][1][1], state.Board[3][0][0]})

	// nothing
	return count
}

func countEqualTo(player FieldState, fields []FieldState) float32 {
	count := float32(0)
	for i := range fields {
		if fields[i] == Empty {
			continue
		} else if fields[i] == player.Flip() {
			return 0
		} else {
			count++
		}
	}
	if count >= 3 {
		return 4
	} else if count >= 2 {
		return 1
	}
	return 0
}

func (state *GameState) GetLegalMoves() []MoveCoordinate {
	moves := make([]MoveCoordinate, 0) // only add items once found (will otherwise initialize to {0,0}, causing errors)
	for x := int8(0); x < 4; x++ {
		for y := int8(0); y < 4; y++ {
			move := MoveCoordinate{x, y}
			if state.findSpot(move) > -1 {
				moves = append(moves, move)
			}
		}
	}
	return moves
}

func (state *GameState) GetString() string {
	outStr := ""
	for z := int8(0); z < 4; z++ {
		for y := int8(0); y < 4; y++ {
			for x := int8(0); x < 4; x++ {
				outStr += state.Board[z][y][x].GetString()
				outStr += " "
			}
			outStr += "\r\n"
		}
		outStr += "\r\n"
	}
	return outStr
}
