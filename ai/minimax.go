package ai

import (
	"cubes/main/engine"
	"fmt"
	"math"
)

type evaluatedState struct {
	move engine.MoveCoordinate
	eval float32
}

var totalPasses int = 0

func GetNextMove(state *engine.GameState, depth int) engine.MoveCoordinate {
	moves := state.GetLegalMoves()
	evaluatedStates := make([]evaluatedState, len(moves))
	totalPasses = 0

	// var wg sync.WaitGroup
	// var mu sync.Mutex

	for i, move := range moves {
		// wg.Add(1)
		func(i int, move engine.MoveCoordinate) {
			// defer wg.Done()
			// mu.Lock()
			evaluatedStates[i] = evaluatedState{moves[i], evaluateState(state, depth, -math.MaxFloat32, math.MaxFloat32, state.CurrentPlayer)}
			// mu.Unlock()
			fmt.Printf("%f:%d,%d\r\n", evaluatedStates[i].eval, evaluatedStates[i].move.X, evaluatedStates[i].move.Y)
		}(i, move)
	}
	// wg.Wait()
	highest := evaluatedStates[0]
	for i := range evaluatedStates {
		// fmt.Printf("eval: %f, move: %d\r\n", evaluatedStates[i].eval, evaluatedStates[i].move)
		if evaluatedStates[i].eval > highest.eval {
			highest = evaluatedStates[i]
		}
	}
	fmt.Printf("Steps taken: %d\r\n", totalPasses)
	// fmt.Scanln()
	return highest.move
}

// RETURNS 1 EVERY TIME?????
func evaluateState(state *engine.GameState, depth int, alpha float32, beta float32, evaluateFor engine.FieldState) float32 {
	totalPasses++
	finished, winner := state.GetWinner()
	if finished {
		if winner == evaluateFor {
			return float32(1_000) + float32(depth)*float32(100)
		} else if winner == evaluateFor.Flip() {
			return float32(-1_000) - float32(depth)*float32(100)
		} else {
			return float32(0)
		}
	}

	// fmt.Printf("%d\r\n", depth)
	legalMoves := state.GetLegalMoves()
	if depth <= 0 || len(legalMoves) == 0 {
		return SimpleEval(state, evaluateFor)
	}

	if evaluateFor == state.CurrentPlayer {
		value := float32(math.SmallestNonzeroFloat32)
		for i := range legalMoves {
			newState := state.GetMovedClone(legalMoves[i])
			value = max(value, evaluateState(newState, depth-1, alpha, beta, evaluateFor.Flip()))
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
			value = min(value, evaluateState(newState, depth-1, alpha, beta, evaluateFor.Flip()))
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
	if returnVal < float32(-100) {
		return -100
	} else if returnVal > float32(100) {
		return 100
	}
	return returnVal
}
