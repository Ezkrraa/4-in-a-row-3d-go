package ai

import (
	"cubes/main/engine"
	"math"
	"sync"
)

type evaluatedState struct {
	move engine.MoveCoordinate
	eval float32
}

func GetNextMove(state *engine.GameState, depth int) engine.MoveCoordinate {
	moves := state.GetLegalMoves()
	evaluatedStates := make([]evaluatedState, len(moves))

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, move := range moves {
		wg.Add(1)
		func(i int, move engine.MoveCoordinate) {
			defer wg.Done()
			mu.Lock()
			evaluatedStates[i] = evaluatedState{move, EvaluateState(state.GetMovedClone(move), depth, -math.MaxFloat32, math.MaxFloat32, state.CurrentPlayer, state.CurrentPlayer)}
			mu.Unlock()
		}(i, move)
	}
	wg.Wait()
	highest := evaluatedStates[0]
	for i := range evaluatedStates {
		if evaluatedStates[i].eval > highest.eval {
			highest = evaluatedStates[i]
		}
	}
	return highest.move
}

// RETURNS 1 EVERY TIME?????
func EvaluateState(state *engine.GameState, depth int, alpha float32, beta float32, evaluateAsPlayer engine.FieldState, maximizingPlayer engine.FieldState) float32 {
	finished, winner := state.GetWinner()
	if !state.IsValid() {
		state.GetString()
		panic("INVALID STATE:\r\n" + state.GetString())
	}

	if finished {
		if winner == evaluateAsPlayer {
			return float32(1_000) + float32(depth)*float32(100)
		} else if winner == evaluateAsPlayer.Flip() {
			return float32(-1_000) - float32(depth)*float32(100)
		} else {
			return float32(0)
		}
	}

	legalMoves := state.GetLegalMoves()
	if depth <= 0 || len(legalMoves) == 0 {
		return SimpleEval(state, evaluateAsPlayer)
	}

	if maximizingPlayer == state.CurrentPlayer {
		value := float32(-math.MaxFloat32)
		for i := range legalMoves {
			newState := state.GetMovedClone(legalMoves[i])
			value = max(value, EvaluateState(newState, depth-1, alpha, beta, evaluateAsPlayer, maximizingPlayer.Flip()))
			if value > beta {
				break
			}
			alpha = max(alpha, value)
		}
		return value
	} else {
		value := float32(math.MaxFloat32)
		for i := range legalMoves {
			newState := state.GetMovedClone(legalMoves[i])
			value = min(value, EvaluateState(newState, depth-1, alpha, beta, evaluateAsPlayer, maximizingPlayer.Flip()))
			if value < alpha {
				break
			}
			beta = min(beta, value)
		}
		return value
	}
}

func SimpleEval(state *engine.GameState, evaluateFor engine.FieldState) float32 {
	good := state.CountNearWins(evaluateFor)
	bad := state.CountNearWins(evaluateFor.Flip())
	if good+bad == 0 {
		return 0
	}
	returnVal := (good - bad) / (good + bad)
	if returnVal < float32(-50) {
		return -50
	} else if returnVal > float32(50) {
		return 50
	}
	return returnVal
}
